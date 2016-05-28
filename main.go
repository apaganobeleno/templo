package main

import (
	"log"
	"net/http"
	"os"
	"templo/routing"

	"github.com/codegangsta/negroni"
)

var port = portWithDefault("3000")

func main() {
	router := routing.BuildRouter()
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	// n.Use(Middleware3)
	n.UseHandler(router)

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
