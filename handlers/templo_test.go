package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CloudyKit/jet"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func init() {
	//For testing proposes
	views = jet.NewHTMLSet("../views")
}

func TestTemplo(t *testing.T) {
	router := muxFor("/hi/{name}", Templo)
	request, _ := http.NewRequest("GET", "/hi/Tony", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)
	assert.Contains(t, string(response.Body.Bytes()), "Hello Tony")
}

func muxFor(path string, handler http.HandlerFunc) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(path, handler).Methods("GET")
	return router
}
