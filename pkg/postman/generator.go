package postman

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Generator generates Postman collections from mux routes
type Generator struct {
	BaseURL        string
	CollectionName string
	Description    string
	Version        string
	bodyGenerator  *BodyGenerator
}

// NewGenerator creates a new Postman collection generator
func NewGenerator(baseURL, collectionName string) *Generator {
	return &Generator{
		BaseURL:        baseURL,
		CollectionName: collectionName,
		Description:    "Auto-generated API collection",
		Version:        "1.0.0",
		bodyGenerator:  NewBodyGenerator(),
	}
}

// Generate creates a Postman collection from a mux router
func (g *Generator) Generate(router *mux.Router) (*Collection, error) {
	collector := NewRouteCollector()
	_ = collector.CollectRoutes(router)
	grouped := collector.GroupByPrefix()

	collection := &Collection{
		Info: Info{
			Name:        g.CollectionName,
			Description: g.Description,
			Schema:      "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
			Version:     g.Version,
		},
	}

	// Add auth folder with collection-level auth using environment variable
	collection.Auth = &Auth{
		Type: "bearer",
		Bearer: []AuthKeyValue{
			{
				Key:   "token",
				Value: "{{token}}",
				Type:  "string",
			},
		},
	}

	// Create folders for each route group
	for prefix, routes := range grouped {
		folder := g.createFolder(prefix, routes)
		collection.Item = append(collection.Item, folder)
	}

	return collection, nil
}

// createFolder creates a folder item for a route group
func (g *Generator) createFolder(prefix string, routes []RouteInfo) Item {
	folder := Item{
		Name:        strings.Title(prefix),
		Description: fmt.Sprintf("Endpoints for %s management", prefix),
		Item:        []Item{},
	}

	for _, route := range routes {
		request := g.createRequest(route)
		folder.Item = append(folder.Item, request)
	}

	return folder
}

// createRequest creates a request item from route info
func (g *Generator) createRequest(route RouteInfo) Item {
	request := &Request{
		Method:      route.Method,
		Header:      g.createHeaders(route),
		URL:         g.createURL(route),
		Description: route.Description,
	}

	// Add body for POST, PUT, PATCH requests
	if route.Method == "POST" || route.Method == "PUT" || route.Method == "PATCH" {
		request.Body = g.createBody(route)
	}

	// Override auth if route doesn't require it
	if !route.RequiresAuth {
		request.Auth = &Auth{
			Type: "noauth",
		}
	}

	item := Item{
		Name:    GenerateName(route),
		Request: request,
	}

	// Add test scripts for common patterns - save token to environment
	if route.Method == "POST" && strings.Contains(route.Path, "/auth/login") {
		item.Event = []Event{
			{
				Listen: "test",
				Script: Script{
					Type: "text/javascript",
					Exec: []string{
						"// Save token to environment variable",
						"var jsonData = pm.response.json();",
						"if (jsonData.accessToken) {",
						"    pm.environment.set('token', jsonData.accessToken);",
						"    console.log('Access token saved to environment');",
						"}",
					},
				},
			},
		}
	}

	return item
}

// createHeaders creates headers for a request
func (g *Generator) createHeaders(route RouteInfo) []Header {
	headers := []Header{
		{
			Key:   "Content-Type",
			Value: "application/json",
		},
	}

	if route.RequiresAuth {
		headers = append(headers, Header{
			Key:         "Authorization",
			Value:       "Bearer {{token}}",
			Description: "JWT authentication token",
		})
	}

	return headers
}

// createURL creates a URL object for a request
func (g *Generator) createURL(route RouteInfo) URL {
	path := route.Path

	// Replace {param} with :param for Postman
	for _, param := range route.PathParams {
		placeholder := fmt.Sprintf("{%s}", param)
		path = strings.ReplaceAll(path, placeholder, fmt.Sprintf(":{{%s}}", param))
	}

	// Split path into parts
	pathParts := []string{}
	for _, part := range strings.Split(strings.Trim(path, "/"), "/") {
		if part != "" {
			pathParts = append(pathParts, part)
		}
	}

	url := URL{
		Raw:      fmt.Sprintf("{{baseUrl}}%s", path),
		Protocol: "http",
		Host:     []string{"{{baseUrl}}"},
		Path:     pathParts,
		Variable: []Variable{},
		Query:    []Query{},
	}

	// Add path variables
	for _, param := range route.PathParams {
		url.Variable = append(url.Variable, Variable{
			Key:         param,
			Value:       fmt.Sprintf("<%s>", param),
			Description: fmt.Sprintf("The %s identifier", param),
		})
	}

	// Add pagination query parameters for GET /article endpoint
	if route.Method == "GET" && route.Path == "/article" {
		url.Query = append(url.Query, Query{
			Key:         "page",
			Value:       "1",
			Description: "Page number (default: 1)",
			Disabled:    true,
		})
		url.Query = append(url.Query, Query{
			Key:         "pageSize",
			Value:       "10",
			Description: "Number of items per page (default: 10, max: 100)",
			Disabled:    true,
		})
	}

	return url
}

// createBody creates a sample body for a request using reflection
func (g *Generator) createBody(route RouteInfo) *Body {
	exampleBody := g.bodyGenerator.GenerateBody(route)

	// Fallback for special cases not covered by struct mapping
	if exampleBody == "{\n  \"key\": \"value\"\n}" {
		exampleBody = g.createFallbackBody(route)
	}

	return &Body{
		Mode: "raw",
		Raw:  exampleBody,
		Options: &BodyOptions{
			Raw: &RawOptions{
				Language: "json",
			},
		},
	}
}

// createFallbackBody handles special cases not in struct models
func (g *Generator) createFallbackBody(route RouteInfo) string {
	path := route.Path

	// Special case handlers
	if strings.Contains(path, "/role") && strings.Contains(path, "/permissions") && route.Method == "POST" {
		return `{
  "roleId": "00000000-0000-0000-0000-000000000000",
  "sectionId": "00000000-0000-0000-0000-000000000000",
  "permissionId": "00000000-0000-0000-0000-000000000000"
}`
	} else if strings.Contains(path, "/user") && strings.Contains(path, "/role") {
		return `{
  "roleId": "00000000-0000-0000-0000-000000000000"
}`
	} else if strings.Contains(path, "/article/rss") {
		return `{
  "rssFeedUrl": "https://example.com/feed.xml"
}`
	} else if strings.Contains(path, "/source") && strings.Contains(path, "/ingest") {
		return `{}`
	}

	// Default fallback
	return `{
  "key": "value"
}`
}

// SaveToFile saves the collection to a JSON file
func (g *Generator) SaveToFile(collection *Collection, filename string) error {
	data, err := json.MarshalIndent(collection, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal collection: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// GenerateWithTimestamp generates a collection and saves it with a timestamp
func (g *Generator) GenerateWithTimestamp(router *mux.Router, outputDir string) (string, error) {
	collection, err := g.Generate(router)
	if err != nil {
		return "", err
	}

	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("%s/%s_%s.postman_collection.json", outputDir,
		strings.ReplaceAll(strings.ToLower(g.CollectionName), " ", "_"), timestamp)

	err = g.SaveToFile(collection, filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}

// GenerateToFile generates and saves collection to a specific file
func (g *Generator) GenerateToFile(router *mux.Router, filename string) error {
	collection, err := g.Generate(router)
	if err != nil {
		return err
	}

	return g.SaveToFile(collection, filename)
}
