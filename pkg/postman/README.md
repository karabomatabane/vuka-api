# Postman Collection Generator

This module automatically generates Postman collections from your Gorilla Mux routes, making it easy to keep your API documentation in sync with your codebase.

## Features

- ✅ Automatically extracts all routes from your mux router
- ✅ Groups endpoints by resource (auth, user, article, etc.)
- ✅ Detects authentication requirements
- ✅ Generates sample request bodies for POST/PUT/PATCH requests
- ✅ Includes path parameters with placeholders
- ✅ Sets up collection variables (baseUrl, token)
- ✅ Auto-saves JWT token from login response
- ✅ Supports both CLI and HTTP endpoint generation

## Usage

### Method 1: CLI Command (Recommended for CI/CD)

Generate the collection using the dedicated CLI tool:

```bash
go run cmd/generate-postman/main.go --generate-postman
```

#### CLI Options

- `--generate-postman`: Generate the collection and exit
- `--output`: Output file path (default: `vuka-api.postman_collection.json`)
- `--base-url`: Base URL for your API (default: `http://localhost:8080`)

#### Examples

```bash
# Generate with default settings
go run cmd/generate-postman/main.go --generate-postman

# Generate with custom output file
go run cmd/generate-postman/main.go --generate-postman --output=docs/api.postman_collection.json

# Generate with production URL
go run cmd/generate-postman/main.go --generate-postman --base-url=https://api.vuka.com
```

### Method 2: HTTP Endpoint (For Development)

Access the collection via HTTP endpoint while your server is running:

```bash
# Download the collection
curl http://localhost:8080/api/postman/collection > vuka-api.postman_collection.json

# Or with custom base URL
curl "http://localhost:8080/api/postman/collection?baseUrl=https://api.vuka.com" > collection.json
```

You can also visit `http://localhost:8080/api/postman/collection` in your browser to download the collection.

### Method 3: Programmatic Usage

Use the generator in your own Go code:

```go
import (
    "vuka-api/pkg/postman"
    "github.com/gorilla/mux"
)

router := mux.NewRouter()
// ... register your routes ...

generator := postman.NewGenerator("http://localhost:8080", "My API")
generator.Description = "Complete API documentation"

// Generate and save to file
err := generator.GenerateToFile(router, "api.postman_collection.json")

// Or generate with timestamp
filename, err := generator.GenerateWithTimestamp(router, "./docs")
```

## Importing into Postman

1. Open Postman
2. Click **Import** button
3. Select the generated `.postman_collection.json` file
4. The collection will be imported with all endpoints organized by resource

## Collection Structure

The generated collection includes:

### Variables
- `baseUrl`: The base URL for all requests (configurable)
- `token`: JWT authentication token (auto-populated on login)

### Folders
Routes are organized into folders by resource:
- **Auth**: Registration and login endpoints
- **User**: User management endpoints
- **Article**: Article CRUD operations
- **Role**: Role management
- **Category**: Category endpoints
- **Directory**: Directory management
- **Source**: RSS source management
- **Newsletter**: Newsletter subscription
- **Permission**: Permission management

### Authentication
- Collection-level Bearer token authentication
- Token automatically saved from login response
- Auth headers included where required

### Sample Requests
Each endpoint includes:
- Appropriate HTTP method
- Path parameters with placeholders
- Sample request bodies (for POST/PUT/PATCH)
- Required headers
- Authentication configuration

## Automation & CI/CD

### Generate on Build

Add to your build script:

```bash
#!/bin/bash
go run cmd/generate-postman/main.go --generate-postman --output=docs/api.postman_collection.json
```

### Git Pre-commit Hook

Create `.git/hooks/pre-commit`:

```bash
#!/bin/bash
go run cmd/generate-postman/main.go --generate-postman
git add *.postman_collection.json
```

### GitHub Actions

```yaml
name: Generate Postman Collection
on:
  push:
    branches: [ main ]

jobs:
  generate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: 1.22
      - name: Generate Collection
        run: go run cmd/generate-postman/main.go --generate-postman
      - name: Commit Collection
        run: |
          git config --local user.email "action@github.com"
          git config --local user.name "GitHub Action"
          git add *.postman_collection.json
          git commit -m "Update Postman collection" || exit 0
          git push
```

## Customization

### Custom Request Bodies

Edit `pkg/postman/generator.go` in the `createBody()` method to customize sample request bodies for your endpoints.

### Authentication Detection

Modify `pkg/postman/collector.go` in `detectAuthMiddleware()` and `detectAdminMiddleware()` to improve authentication detection based on your middleware patterns.

### Route Naming

Customize route names in `GenerateName()` function in `pkg/postman/collector.go`.

## Architecture

The module consists of four main components:

1. **types.go**: Postman Collection v2.1 schema definitions
2. **collector.go**: Extracts and analyzes routes from mux router
3. **generator.go**: Builds Postman collection from route information
4. **CLI/HTTP interfaces**: Multiple ways to trigger generation

## Troubleshooting

### Collection is empty or missing routes

Ensure all routes are registered before generating the collection. The generator walks the entire router tree.

### Authentication not detected correctly

The module uses heuristics to detect auth requirements. For more accurate detection, you may need to customize the `detectAuthMiddleware()` function to match your middleware patterns.

### Path parameters not working

The generator converts `{param}` style parameters to Postman's `:{{param}}` format. If you use a different style, update the `createURL()` function.

## Benefits

- **Stay in Sync**: Collection always matches your codebase
- **Collaborate Easily**: Share up-to-date collection with frontend developers
- **Version Control**: Track API changes through collection diffs
- **Automate Testing**: Use collection for automated API testing
- **Onboarding**: New team members get instant API documentation

## Future Enhancements

Potential improvements:
- Response examples from actual API responses
- Request validation based on struct tags
- OpenAPI/Swagger export
- Environment file generation
- Integration with Postman API for auto-sync
