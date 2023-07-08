package main

import (
	"api-go-gin/database"
	"api-go-gin/router"
)

func main() {
	database.ConnectDatabase()
	router.HandleRequests()
}
