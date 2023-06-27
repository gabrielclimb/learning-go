package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"main.go/models"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := ConnectDB()
	defer db.Close()
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("error")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	product := []models.Product{
		{"Camiseta", "Azul clara", 39, 3},
		{"tenis", "Comfortable", 150, 1},
		{"Fone Bluetooth", "Fone sem fio", 109, 10},
		{"Notebook", "A Fast PC", 1200, 2},
	}

	err := templ.ExecuteTemplate(w, "Index", product)

	if err != nil {
		fmt.Println("error")
	}
}

func ConnectDB() *sql.DB {
	connection := "user=postgres dbname=store password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
