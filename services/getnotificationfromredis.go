package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"realtime-notifications-service/redisclient"
)

type Response struct {
	Message string
}

func GetNotificationFromRedis(userId string) *[]NotificationSchema {
	key := fmt.Sprintf("notifications: %s", userId)
	client := redisclient.GetRedisClient()
	notificationJSONs, err := client.LRange(context.Background(), key, 0, -1).Result()

	if err != nil {
		log.Println("Error while reading the notfications from redis")
	}

	var userNotifications []NotificationSchema
	for _, notificationJSON := range notificationJSONs {
		var notification NotificationSchema
		if err := json.Unmarshal([]byte(notificationJSON), &notification); err != nil {
			log.Println("Error while unmarshaling the notifications from redis", err.Error())
		}
		userNotifications = append(userNotifications, notification)
	}

	return &userNotifications
}
