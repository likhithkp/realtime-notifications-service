package main

import (
	"log"
	"net/http"
	"realtime-notifications-service/config"
	"realtime-notifications-service/handler"
	"realtime-notifications-service/redisclient"
	"realtime-notifications-service/services"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	config.ConnectDB()
	defer config.CloseDB()

	redisClient := redisclient.GetRedisClient()
	defer redisClient.Close()

	go services.ListenNotificationEvents("localhost:9092", "notificationsGroup", "notifications")

	http.HandleFunc("POST /createUser", handler.CreateUser)
	http.HandleFunc("POST /createNotification", handler.ProduceNotification)
	http.HandleFunc("GET /notifications/{user_id}", handler.GetNotifications)
	http.HandleFunc("/ws/{user_id}", handler.GetLiveNotifications)
	http.ListenAndServe(":3000", nil)
}
