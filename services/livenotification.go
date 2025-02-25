package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

func SendLiveNotification(notification *Notification) (sent bool) {
	notificationJson, err := json.Marshal(&notification)
	if err != nil {
		log.Println("Error while marshaling notification for WS", err.Error())
		return
	}

	url := fmt.Sprintf("ws://localhost:3000/ws/%s", strconv.Itoa(notification.UserId))

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	fmt.Println("Connected live to Notification service")

	err = conn.WriteMessage(websocket.TextMessage, []byte(notificationJson))
	if err != nil {
		log.Println("Write error:", err)
	}

	_, response, err := conn.ReadMessage()
	if err != nil {
		log.Println("Read error:", err)
	}

	fmt.Println("Received:", string(response))
	sent = true
	return sent
}
