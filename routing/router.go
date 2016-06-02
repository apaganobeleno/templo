package routing

import (
	"net/http"

	"github.com/apaganobeleno/templo/handlers"
	"github.com/bmizerany/pat"
)

// BuildRouter returns the router we're going to use to handle our app,
// usually in combination with Negroni.
func BuildRouter() *pat.PatternServeMux {
	router := pat.New()
	router.Get("/", http.HandlerFunc(handlers.Welcome))
	router.Get("/hi/:name", http.HandlerFunc(handlers.Templo))
	return router
}
