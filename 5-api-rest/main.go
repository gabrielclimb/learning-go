package main

import (
	"api-rest/database"
	"api-rest/router"
	"fmt"
)

func main() {
	database.ConnectDatabase()
	fmt.Println("Iniciando o servidor Rest com Go")
	router.HandleRequest()
}
