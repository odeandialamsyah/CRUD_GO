package routes

import (
	"net/http"

	"github.com/odeandialamsyah/crud_go/controller"
)

func MapRoutes(server *http.ServeMux) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/employee", controller.NewIndexController())
}