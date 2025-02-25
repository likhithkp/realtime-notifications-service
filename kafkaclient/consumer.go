package kafkaclient

import (
	"log"
	"os"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	consumer *kafka.Consumer
	syncOnce sync.Once
)

func GetConsumer(host string, groupId string) *kafka.Consumer {
	syncOnce.Do(func() {
		c, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": host,
			"group.id":          groupId,
			"auto.offset.reset": "earliest"})

		if err != nil {
			log.Printf("Failed to create consumer: %s", err)
			os.Exit(1)
		}
		consumer = c
	})
	return consumer
}
