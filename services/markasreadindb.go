package services

import (
	"context"
	"log"
	"realtime-notifications-service/config"
)

func MarkAsReadInDb(userId string, notificationId string) (value int64) {
	query := `UPDATE notifications SET is_read=$1 WHERE id=$2 AND user_id=$3`

	res, err := config.DB.Exec(context.Background(), query, true, notificationId, userId)
	if err != nil {
		log.Println("Error while marking notification as read", err.Error())
		return
	}
	return res.RowsAffected()
}
