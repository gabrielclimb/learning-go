package main

import (
	"fmt"
	"net/http"
	"text/template"

	"main.go/internals/adapters/database"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("error")
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	db := database.ConnectDB()
	defer db.DB.Close()

	products, err := db.GetAllProducts()

	err = templ.ExecuteTemplate(w, "Index", products)

	if err != nil {
		fmt.Println("error")
	}
}
