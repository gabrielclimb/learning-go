package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"main.go/models"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts, err := models.GetAllProducts()
	err = templ.ExecuteTemplate(w, "Index", allProducts)
	if err != nil {
		fmt.Println("error")
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "New", nil)
}
