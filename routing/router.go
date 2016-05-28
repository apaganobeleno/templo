package routing

import (
	"net/http"
	"templo/handlers"

	"github.com/gorilla/mux"
)

func BuildRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/hi/{name}", handlers.Templo).Methods("GET")
	router.PathPrefix("/assets/").Handler(http.FileServer(http.Dir("./public/")))
	router.Handle("/", router)

	return router
}
