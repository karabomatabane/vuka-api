# Vuka API

Backend API for the Vuka platform.

## Quick Start

### Prerequisites
- Go 1.22+
- PostgreSQL
- Docker (optional)

### Running the API

```bash
# Install dependencies
go mod download

# Run the application
go run cmd/main.go
```

## Postman Collection

This project includes automatic Postman collection generation to make API collaboration easier.

### Generate Collection

```bash
# Using Make (recommended)
make postman

# Or using Go directly
go run cmd/generate-postman/main.go --generate-postman

# For production URLs
make postman-prod
```

### Download via HTTP

While the server is running:

```bash
curl http://localhost:3000/api/postman/collection > vuka-api.postman_collection.json
```

Or visit `http://localhost:3000/api/postman/collection` in your browser.

### Import to Postman

1. Open Postman
2. Click **Import**
3. Select the generated `vuka-api.postman_collection.json` file
4. Start testing your API!

The collection includes:
- ✅ All API endpoints organized by resource
- ✅ Sample request bodies
- ✅ Authentication setup with auto-token saving
- ✅ Environment variables (baseUrl, token)

For detailed documentation, see [pkg/postman/README.md](pkg/postman/README.md)

## Project Structure

```
vuka-api/
├── cmd/
│   ├── main.go                    # Main application entry
│   └── generate-postman/          # Postman collection generator CLI
├── pkg/
│   ├── config/                    # Configuration
│   ├── controllers/               # HTTP handlers
│   ├── middleware/                # HTTP middleware
│   ├── models/                    # Data models
│   ├── postman/                   # Postman collection generator
│   ├── repository/                # Data access layer
│   ├── routes/                    # Route definitions
│   ├── services/                  # Business logic
│   └── utils/                     # Utilities
└── vuka-cms/                      # Frontend application
```

## Development

### Available Make Commands

```bash
make help          # Show all available commands
make build         # Build the application
make run           # Run the application
make test          # Run tests
make postman       # Generate Postman collection
make clean         # Clean build artifacts
```

## API Documentation

The API documentation is automatically maintained through the Postman collection. Generate the latest collection anytime with `make postman` and share it with your team.

## Contributing

1. Fork the repository
2. Create your feature branch
3. Make your changes
4. Generate updated Postman collection: `make postman`
5. Commit your changes (include the updated collection)
6. Push to the branch
7. Create a Pull Request

## License

[NO LICENCE AS YET]
