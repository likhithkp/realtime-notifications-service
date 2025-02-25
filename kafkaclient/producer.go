package kafkaclient

import (
	"log"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	producer *kafka.Producer
	once     sync.Once
)

func GetProducer(host string) *kafka.Producer {
	once.Do(func() {
		p, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": host,
			"acks":              "all",
		})

		if err != nil {
			log.Fatalf("Unable to create Kafka producer: %s", err)
		}

		go func() {
			for e := range p.Events() {
				switch ev := e.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						log.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
					} else {
						log.Printf("Produced event to topic %s: key = %s value = %s\n",
							*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
					}
				case *kafka.Error:
					log.Println(ev.Error())
				}
			}
		}()

		producer = p
	})

	return producer
}
