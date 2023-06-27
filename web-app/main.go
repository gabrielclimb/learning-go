package main

import (
	"fmt"
	"net/http"

	"main.go/internals/adapters/view"
)

func main() {
	view.LoadRoutes()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("error")
	}
}
