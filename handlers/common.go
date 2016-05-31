package handlers

import (
	"os"

	"github.com/CloudyKit/jet"
)

var views = jet.NewHTMLSet("./views")

func init() {
	if os.Getenv("DEVELOPMENT") == "1" {
		views.SetDevelopmentMode(true) //Builds templates on the flight
	}
}
