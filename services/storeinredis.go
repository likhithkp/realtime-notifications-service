package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"realtime-notifications-service/redisclient"
	"strconv"
	"time"
)

type NotificationSchema struct {
	ID       string         `json:"id"`
	UserId   string         `json:"user_id"`
	Event    string         `json:"event"`
	Message  string         `json:"message"`
	Metadata map[string]any `json:"metadata"`
}

var ctx = context.Background()

func StoreNotificationRedis(notification *Notification, notificationId string) {
	newNotification := NotificationSchema{
		ID:       notificationId,
		UserId:   strconv.Itoa(notification.UserId),
		Event:    notification.Event,
		Message:  notification.Message,
		Metadata: notification.Metadata,
	}

	newNotificationJson, err := json.Marshal(newNotification)
	if err != nil {
		log.Println("Error while marshaling new notification to JSON", err.Error())
		return
	}

	key := fmt.Sprintf("notifications: %s", newNotification.UserId)

	client := redisclient.GetRedisClient()
	client.LPush(ctx, key, newNotificationJson)

	exists, err := client.Exists(ctx, key).Result()

	if err != nil {
		log.Println("Error while checking is the key already exists in redis", err.Error())
		return
	}

	if exists == 0 {
		client.Expire(ctx, key, time.Hour*24*7)
	}
}
