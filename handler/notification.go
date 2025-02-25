package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"realtime-notifications-service/services"
)

type Notification struct {
	UserId   int            `json:"user_id"`
	Event    string         `json:"event"`
	Message  string         `json:"message"`
	Metadata map[string]any `json:"metadata"`
}

func ProduceNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Not a valid method", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	notification := new(Notification)

	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	byteData, err := json.Marshal(&notification)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while marshalling: %v\n", err)
		return
	}

	err = services.PublishNotificationEvent("notifications", "notification", byteData, "localhost:9092")
	if err != nil {
		log.Printf("Failed to publish event: %v\n", err)
		http.Error(w, "Failed to publish notification", http.StatusInternalServerError)
		return
	}

	log.Println("✅ Event published successfully!")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Notification sent successfully!"})
}
