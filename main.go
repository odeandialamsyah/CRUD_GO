package main

import (
	"net/http"

	"github.com/odeandialamsyah/crud_go/database"
	"github.com/odeandialamsyah/crud_go/routes"
)

func main() {
	database.InitDatabase()
	
	server := http.NewServeMux()

	routes.MapRoutes(server)

	http.ListenAndServe(":8080", server)
}