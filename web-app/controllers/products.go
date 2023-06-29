package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

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

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			log.Println("Price Conversion Error:", err.Error())
		}
		amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
		if err != nil {
			log.Println("Amount Conversion Error:", err.Error())
		}

		models.AddNewProduct
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
