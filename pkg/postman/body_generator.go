package postman

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"vuka-api/pkg/models"
	"vuka-api/pkg/models/db"

	"github.com/google/uuid"
)

// BodyGenerator generates sample request bodies from Go structs
type BodyGenerator struct {
	modelMap map[string]interface{}
}

// NewBodyGenerator creates a new body generator with model mappings
func NewBodyGenerator() *BodyGenerator {
	bg := &BodyGenerator{
		modelMap: make(map[string]interface{}),
	}
	bg.registerModels()
	return bg
}

// registerModels registers all API models for body generation
func (bg *BodyGenerator) registerModels() {
	// Auth models
	bg.modelMap["/auth/register"] = models.RegisterBody{}
	bg.modelMap["/auth/login"] = models.LoginBody{}

	// Article models
	bg.modelMap["/article_POST"] = db.Article{}
	bg.modelMap["/article_PATCH"] = db.Article{}
	bg.modelMap["/article_PUT"] = db.Article{}

	// User models
	bg.modelMap["/user_PATCH"] = db.User{}

	// Role models
	bg.modelMap["/role_POST"] = db.Role{}
	bg.modelMap["/role_PATCH"] = db.Role{}

	// Source models
	bg.modelMap["/source_POST"] = db.Source{}
	bg.modelMap["/source_PATCH"] = db.Source{}

	// Category models
	bg.modelMap["/category_POST"] = db.Category{}
	bg.modelMap["/category_PATCH"] = db.Category{}

	// Directory models
	bg.modelMap["/directory_POST"] = db.DirectoryCategory{}
	bg.modelMap["/directory/entries_POST"] = db.DirectoryEntry{}

	// Newsletter models
	bg.modelMap["/newsletter/subscribe"] = db.NewsletterSubscriber{}
	bg.modelMap["/newsletter_PATCH"] = db.NewsletterSubscriber{}

	// Permission models
	bg.modelMap["/permission_POST"] = db.Permission{}
	bg.modelMap["/permission_PATCH"] = db.Permission{}

	// Section models
	bg.modelMap["/section_POST"] = db.Section{}
	bg.modelMap["/section_PATCH"] = db.Section{}
}

// GenerateBody generates a sample JSON body for a route
func (bg *BodyGenerator) GenerateBody(route RouteInfo) string {
	// Try to find exact path match first
	if model, ok := bg.modelMap[route.Path]; ok {
		return bg.structToJSON(model)
	}

	// Try path + method combination
	key := route.Path + "_" + route.Method
	if model, ok := bg.modelMap[key]; ok {
		return bg.structToJSON(model)
	}

	// Try to find by path prefix
	for path, model := range bg.modelMap {
		if strings.HasPrefix(route.Path, strings.Split(path, "_")[0]) {
			if strings.Contains(path, "_") {
				// Path with method
				parts := strings.Split(path, "_")
				if parts[1] == route.Method {
					return bg.structToJSON(model)
				}
			}
		}
	}

	// Fallback
	return `{
  "key": "value"
}`
}

// structToJSON converts a struct to sample JSON with example values
func (bg *BodyGenerator) structToJSON(model interface{}) string {
	v := reflect.TypeOf(model)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	result := make(map[string]interface{})

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Get JSON tag
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// Parse JSON tag (handle omitempty, etc.)
		jsonName := strings.Split(jsonTag, ",")[0]
		if jsonName == "" {
			jsonName = field.Name
		}

		// Skip embedded structs and relations
		if field.Anonymous ||
			field.Name == "Model" ||
			field.Name == "CreatedAt" ||
			field.Name == "UpdatedAt" ||
			field.Name == "DeletedAt" ||
			field.Name == "ID" {
			continue
		}

		// Skip relationship fields (slices and complex types for POST/PATCH)
		if field.Type.Kind() == reflect.Slice ||
			field.Type.Kind() == reflect.Array ||
			(field.Type.Kind() == reflect.Struct && field.Type.Name() != "UUID" && field.Type.Name() != "Time") {
			continue
		}

		// Generate example value based on type
		result[jsonName] = bg.generateExampleValue(field)
	}

	// Pretty print JSON
	jsonBytes, _ := json.MarshalIndent(result, "", "  ")
	return string(jsonBytes)
}

// generateExampleValue generates an example value for a field
func (bg *BodyGenerator) generateExampleValue(field reflect.StructField) interface{} {
	fieldName := strings.ToLower(field.Name)

	switch field.Type.Kind() {
	case reflect.String:
		// Special cases based on field name
		if strings.Contains(fieldName, "email") {
			return "user@example.com"
		} else if strings.Contains(fieldName, "password") {
			return "SecurePassword123!"
		} else if strings.Contains(fieldName, "url") {
			return "https://example.com"
		} else if strings.Contains(fieldName, "phone") {
			return "+1234567890"
		} else if strings.Contains(fieldName, "name") || strings.Contains(fieldName, "username") {
			return "Example Name"
		} else if strings.Contains(fieldName, "title") {
			return "Example Title"
		} else if strings.Contains(fieldName, "description") || strings.Contains(fieldName, "summary") {
			return "Example description"
		} else if strings.Contains(fieldName, "content") || strings.Contains(fieldName, "body") {
			return "Example content"
		} else if strings.Contains(fieldName, "language") {
			return "en"
		} else if strings.Contains(fieldName, "type") {
			return "default"
		}
		return "example value"

	case reflect.Bool:
		return false

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return 0

	case reflect.Float32, reflect.Float64:
		return 0.0

	case reflect.Ptr:
		// Handle pointer types
		elemType := field.Type.Elem()
		if elemType.Name() == "UUID" || strings.Contains(fieldName, "id") {
			return "00000000-0000-0000-0000-000000000000"
		}
		return nil

	case reflect.Struct:
		// Handle specific struct types
		if field.Type == reflect.TypeOf(uuid.UUID{}) {
			return "00000000-0000-0000-0000-000000000000"
		}
		if field.Type == reflect.TypeOf(time.Time{}) {
			return time.Now().Format(time.RFC3339)
		}
		return nil

	default:
		return nil
	}
}
