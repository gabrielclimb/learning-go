package view

import (
	"net/http"

	"main.go/internals/adapters/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
