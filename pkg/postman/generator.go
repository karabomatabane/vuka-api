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
}

// NewGenerator creates a new Postman collection generator
func NewGenerator(baseURL, collectionName string) *Generator {
	return &Generator{
		BaseURL:        baseURL,
		CollectionName: collectionName,
		Description:    "Auto-generated API collection",
		Version:        "1.0.0",
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
		Variable: []Variable{
			{
				Key:   "baseUrl",
				Value: g.BaseURL,
				Type:  "string",
			},
			{
				Key:         "token",
				Value:       "",
				Type:        "string",
				Description: "JWT authentication token",
			},
		},
	}

	// Add auth folder with collection-level auth
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

	// Add test scripts for common patterns
	if route.Method == "POST" && strings.Contains(route.Path, "/auth/login") {
		item.Event = []Event{
			{
				Listen: "test",
				Script: Script{
					Type: "text/javascript",
					Exec: []string{
						"var jsonData = pm.response.json();",
						"if (jsonData.token) {",
						"    pm.collectionVariables.set('token', jsonData.token);",
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
	}

	// Add path variables
	for _, param := range route.PathParams {
		url.Variable = append(url.Variable, Variable{
			Key:         param,
			Value:       fmt.Sprintf("<%s>", param),
			Description: fmt.Sprintf("The %s identifier", param),
		})
	}

	return url
}

// createBody creates a sample body for a request
func (g *Generator) createBody(route RouteInfo) *Body {
	var exampleBody string

	// Generate example bodies based on path
	if strings.Contains(route.Path, "/auth/register") {
		exampleBody = `{
  "email": "user@example.com",
  "password": "SecurePassword123!",
  "name": "John Doe"
}`
	} else if strings.Contains(route.Path, "/auth/login") {
		exampleBody = `{
  "email": "user@example.com",
  "password": "SecurePassword123!"
}`
	} else if strings.Contains(route.Path, "/article") {
		if route.Method == "POST" {
			exampleBody = `{
  "title": "Article Title",
  "content": "Article content goes here",
  "categoryId": "category-uuid"
}`
		} else if route.Method == "PUT" || route.Method == "PATCH" {
			exampleBody = `{
  "title": "Updated Article Title",
  "content": "Updated content"
}`
		}
	} else if strings.Contains(route.Path, "/user") && route.Method == "PATCH" {
		exampleBody = `{
  "name": "Updated Name",
  "email": "updated@example.com"
}`
	} else if strings.Contains(route.Path, "/role") {
		if route.Method == "POST" {
			exampleBody = `{
  "name": "Role Name",
  "description": "Role description"
}`
		} else if route.Method == "PATCH" {
			exampleBody = `{
  "name": "Updated Role Name"
}`
		}
	} else {
		// Generic body
		exampleBody = `{
  "key": "value"
}`
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
