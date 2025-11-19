.PHONY: help postman postman-prod build run test clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

postman: ## Generate Postman collection for local development
	@echo "Generating Postman collection..."
	@go run cmd/generate-postman/main.go --generate-postman --output=vuka-api.postman_collection.json --base-url=http://localhost:8080
	@echo "Collection generated: vuka-api.postman_collection.json"

postman-prod: ## Generate Postman collection for production
	@echo "Generating Postman collection for production..."
	@go run cmd/generate-postman/main.go --generate-postman --output=vuka-api-prod.postman_collection.json --base-url=https://api.vuka.com
	@echo "Collection generated: vuka-api-prod.postman_collection.json"

build: ## Build the application
	@echo "Building application..."
	@go build -o bin/vuka-api cmd/main.go
	@echo "Build complete: bin/vuka-api"

run: ## Run the application
	@go run cmd/main.go

test: ## Run tests
	@go test ./... -v

clean: ## Clean build artifacts
	@rm -rf bin/
	@rm -f *.postman_collection.json
	@echo "Cleaned build artifacts"
