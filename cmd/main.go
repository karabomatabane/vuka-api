package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"vuka-api/pkg/config"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/routes"
	"vuka-api/pkg/services"

	"github.com/gorilla/mux"
)

func init() {
	config.LoadEnvVariables()
	config.Connect()
}

func MigrateSources(service *services.SourceService, csvPath string) {
	file, err := os.Open(csvPath)
	if err != nil {
		log.Fatalf("failed to open sources.csv: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read sources.csv: %v", err)
	}

	for i, row := range records {
		if i == 0 {
			continue // skip header
		}
		name := row[0]
		website := row[1]
		rss := row[2]
		source := db.Source{
			Name:       name,
			WebsiteUrl: website,
			RssFeedUrl: rss,
		}
		if err := service.CreateSource(&source); err != nil {
			log.Printf("failed to insert source %s: %v", name, err)
		}
	}
	log.Println("Source migration complete.")
}

func main() {
	router := mux.NewRouter()
	routes.RegisterAuthRoutes(router)
	routes.RegisterUserRoutes(router)
	routes.RegisterArticleRoutes(router)
	routes.RegisterRoleRoutes(router)
	routes.RegisterSourceRoutes(router)

	// Migrate sources from CSV on startup
	serviceManager := services.NewServices(config.GetDB())
	MigrateSources(serviceManager.Source, "bin/sources.csv")

	http.Handle("/", router)
	listeningAddr := fmt.Sprintf("localhost:%v", os.Getenv("PORT"))
	log.Printf("Server is running on %s", listeningAddr)
	log.Fatal(http.ListenAndServe(listeningAddr, router))
}
