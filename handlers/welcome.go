package handlers

import (
	"log"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	view, err := views.GetTemplate("welcome.jet")
	if err != nil {
		log.Println("Unexpected template err:", err.Error())
		return
	}

	view.Execute(w, nil, nil)
}
