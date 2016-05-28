package routing

import (
	"net/http"
	"templo/handlers"

	"github.com/bmizerany/pat"
)

func BuildRouter() *pat.PatternServeMux {
	router := pat.New()
	router.Get("/hi/:name", http.HandlerFunc(handlers.Templo))
	return router
}
