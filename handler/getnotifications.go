package handler

import (
	"encoding/json"
	"net/http"
	"realtime-notifications-service/services"
)

func GetNotifications(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Not a valid method", http.StatusMethodNotAllowed)
		return
	}

	userId := r.PathValue("user_id")
	if userId == "" {
		http.Error(w, "Enter the user id", http.StatusBadRequest)
		return
	}

	notifications := services.GetNotificationFromRedis(userId)
	if notifications != nil {
		json.NewEncoder(w).Encode(&notifications)
		return
	}

	// If Redis is empty, fetch from DB
	notificationsFromDb := services.GetNotificationsFromDB(userId)
	json.NewEncoder(w).Encode(&notificationsFromDb)

}
