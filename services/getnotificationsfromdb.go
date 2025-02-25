package services

import (
	"context"
	"log"
	"realtime-notifications-service/config"
)

func GetNotificationsFromDB(userId string) *[]NotificationSchema {
	query := `SELECT id, user_id, event, message, metadata FROM notifications WHERE user_id = $1 AND is_read = FALSE ORDER BY created_at DESC`

	rows, err := config.DB.Query(context.Background(), query, userId)
	if err != nil {
		log.Fatal("Error while getting data from row", err.Error())
	}

	defer rows.Close()

	var notifications []NotificationSchema
	for rows.Next() {
		var notification NotificationSchema
		err = rows.Scan(&notification.ID, &notification.UserId, &notification.Event, &notification.Message, &notification.Metadata)
		if err != nil {
			log.Fatal("Error while scanning data from row", err.Error())
		}
		notifications = append(notifications, notification)
	}

	return &notifications
}
