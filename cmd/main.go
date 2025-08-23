package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"vuka-api/pkg/config"
	"vuka-api/pkg/routes"

	"github.com/gorilla/mux"
)

func init() {
	config.LoadEnvVariables()
	config.Connect()
}

func main() {
	router := mux.NewRouter()
	routes.RegisterAuthRoutes(router)
	routes.RegisterUserRoutes(router)
	routes.RegisterFilmRoutes(router)
	http.Handle("/", router)
	listeningAddr := fmt.Sprintf("localhost:%v", os.Getenv("PORT"))
	log.Printf("Server is running on %s", listeningAddr)
	log.Fatal(http.ListenAndServe(listeningAddr, router))
}
