package main

import (
	"messenger/app/models"
)

func main() {
	models.ConnectDatabase()
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run()
}
