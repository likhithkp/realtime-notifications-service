package services

import (
	"context"
	"log"
	"realtime-notifications-service/config"
)

type ResponseMessage struct {
	Message string
	UserId  *string
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
		UserId:  &user_id,
	}

	return &res

}
