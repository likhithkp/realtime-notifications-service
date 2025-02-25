package services

import (
	"log"
	"os"
	"os/signal"
	"realtime-notifications-service/kafkaclient"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ListenNotificationEvents(host string, groupId string, topic string) {
	c := kafkaclient.GetConsumer(host, groupId)
	defer c.Close()

	err := c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Printf("Error while subscribing :%v", err.Error())
		return
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case sig := <-sigChan:
			log.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(1 * time.Second)
			if err != nil {
				if err.(kafka.Error).Code() == kafka.ErrTimedOut {
					continue
				}
				log.Printf("Consumer error: %v\n", err)
				continue
			}

			log.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
			CreateNotificationService(ev.Value)
		}
	}
}
