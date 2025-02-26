package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"realtime-notifications-service/redisclient"
	"realtime-notifications-service/services"
)

type Response struct {
	Message   any
	StatuCode int
}

func MarkNotificationAsRead(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	userId := r.PathValue("user_id")
	notifcationId := r.PathValue("notification_id")

	if userId == "" || notifcationId == "" {
		http.Error(w, "user id and notification id are required", http.StatusBadRequest)
		return
	}

	key := fmt.Sprintf("notifications: %s", userId)

	client := redisclient.GetRedisClient()
	notifications, err := client.LRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		log.Println("Error while fetching notifications from redis", err.Error())
		return
	}

	for _, notification := range notifications {
		var data map[string]any
		if err := json.Unmarshal([]byte(notification), &data); err != nil {
			log.Println("Error while unmarshaling the notification")
			return
		}

		if fmt.Sprintf("%v", data["id"]) == notifcationId {
			err := client.LRem(context.Background(), key, 1, notification).Err()
			if err != nil {
				log.Println("Error while removing the data from list", err.Error())
				http.Error(w, "Error while marking the notfication as read", http.StatusBadGateway)
				return
			}

			value := services.MarkAsReadInDb(userId, notifcationId)

			if value == 1 {
				res := Response{
					Message:   "Marked as read",
					StatuCode: 200,
				}

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(&res)
			}
		}
	}

}
