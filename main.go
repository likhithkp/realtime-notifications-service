package main

import (
	"log"
	"net/http"
	"realtime-notifications-service/config"
	"realtime-notifications-service/handler"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	config.ConnectDB()
	defer config.CloseDB()

	http.HandleFunc("/createUser", handler.CreateUser)
	http.HandleFunc("/createNotification", handler.ProduceNotification)
	http.ListenAndServe(":3000", nil)
}
