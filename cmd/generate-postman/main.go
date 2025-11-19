package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"vuka-api/pkg/config"
	"vuka-api/pkg/postman"
	"vuka-api/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Define CLI flags
	generatePostman := flag.Bool("generate-postman", false, "Generate Postman collection and exit")
	outputFile := flag.String("output", "vuka-api.postman_collection.json", "Output file for Postman collection")
	baseURL := flag.String("base-url", "http://localhost:8080", "Base URL for API")
	flag.Parse()

	// Load environment variables
	config.LoadEnvVariables()

	// Create router and register routes
	router := mux.NewRouter()
	routes.RegisterAuthRoutes(router)
	routes.RegisterUserRoutes(router)
	routes.RegisterArticleRoutes(router)
	routes.RegisterRoleRoutes(router)
	routes.RegisterSourceRoutes(router)
	routes.RegisterCategoryRoutes(router)
	routes.RegisterDirectoryRoutes(router)
	routes.RegisterPermissionRoutes(router)
	routes.RegisterNewsletterRoutes(router)

	if *generatePostman {
		// Generate Postman collection
		generator := postman.NewGenerator(*baseURL, "Vuka API")
		generator.Description = "Complete API documentation for Vuka API"

		err := generator.GenerateToFile(router, *outputFile)
		if err != nil {
			log.Fatalf("Failed to generate Postman collection: %v", err)
		}

		fmt.Printf("✅ Postman collection generated successfully: %s\n", *outputFile)
		os.Exit(0)
	}

	log.Println("Use --generate-postman flag to generate Postman collection")
	log.Println("Example: go run cmd/generate-postman/main.go --generate-postman --output=collection.json")
}
