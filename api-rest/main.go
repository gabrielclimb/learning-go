package main

import (
	"api-rest/models"
	"api-rest/router"
	"fmt"
)

func main() {
	models.Personalities = []models.Personality{
		{ID: 1, Name: "Nome 1", History: "Historia 1"},
		{ID: 2, Name: "Nome 2", History: "Historia 2"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	router.HandleRequest()
}
