package main

import (
	"fmt"
	"net/http"
	"text/template"
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
	err := templ.ExecuteTemplate(w, "Index", nil)

	if err != nil {
		fmt.Println("error")
	}
}
