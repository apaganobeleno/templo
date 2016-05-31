package main

import (
	"log"
	"net/http"
	"os"
	"templo/routing"

	"github.com/codegangsta/negroni"
	"github.com/joho/godotenv"
)

var port = portWithDefault("3000")

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("| WARNING: Could not find .env file relaying on system ENV")
	}

	router := routing.BuildRouter()
	n := negroni.Classic()
	n.UseHandler(router)

	if os.Getenv("DEVELOPMENT") == "1" {
		log.Println("| WARNING: App running in development mode")
	}

	log.Printf("| App running on port %v", port)
	http.ListenAndServe(":"+port, n)
}

func portWithDefault(other string) string {
	port := os.Getenv("PORT")
	if port == "" {
		return other
	}
	return port
}
