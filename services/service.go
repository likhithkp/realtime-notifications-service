package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"realtime-notifications-service/config"
)

type ResponseMessage struct {
	Message string
	Id      *string
}
type Notification struct {
	UserId   int            `json:"user_id"`
	Event    string         `json:"event"`
	Message  string         `json:"message"`
	Metadata map[string]any `json:"metadata"`
}

func CreateUserService(name string, email string, password string) (message *ResponseMessage) {
	var user_id string
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`

	if err := config.DB.QueryRow(context.Background(), query, name, email, password).Scan(&user_id); err != nil {
		log.Println("Error while inserting into users", err.Error())
		return
	}

	res := ResponseMessage{
		Message: "User added",
		Id:      &user_id,
	}

	return &res

}

func CreateNotificationService(eventData []byte) error {
	notification := new(Notification)

	if err := json.Unmarshal(eventData, notification); err != nil {
		log.Printf("Error unmarshaling notification event: %v", err)
		return err
	}

	if notification.UserId == 0 || notification.Event == "" || notification.Message == "" {
		log.Println("Invalid notification data: missing required fields")
		return fmt.Errorf("invalid notification data")
	}

	sent := SendLiveNotification(notification)
	if sent {
		query := `INSERT INTO notifications (user_id, event, message, metadata, is_read) VALUES ($1, $2, $3, $4, $5) RETURNING id`

		var notificationID string
		err := config.DB.QueryRow(context.Background(), query, notification.UserId, notification.Event, notification.Message, notification.Metadata, true).Scan(&notificationID)
		if err != nil {
			log.Printf("Error inserting into read notifications: %v", err)
			return err
		}

		log.Printf("Read notification successfully inserted with ID: %s", notificationID)
		return nil
	}

	query := `INSERT INTO notifications (user_id, event, message, metadata) VALUES ($1, $2, $3, $4) RETURNING id`

	var notificationID string
	err := config.DB.QueryRow(context.Background(), query, notification.UserId, notification.Event, notification.Message, notification.Metadata).Scan(&notificationID)
	if err != nil {
		log.Printf("Error inserting into notifications: %v", err)
		return err
	}

	log.Printf("Notification successfully inserted with ID: %s", notificationID)
	StoreNotificationRedis(notification, notificationID)
	return nil
}
