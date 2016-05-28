package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet"
	"github.com/gorilla/mux"
)

var views = jet.NewHTMLSet("./views")

//Templo is a sample handler to show off some of the components of the template
func Templo(w http.ResponseWriter, r *http.Request) {
	view, err := views.GetTemplate("templo.jet")
	vars := mux.Vars(r)

	if err != nil {
		log.Println("Unexpected template err:", err.Error())
	}
	view.Execute(w, nil, struct {
		Name string
	}{vars["name"]})
}
