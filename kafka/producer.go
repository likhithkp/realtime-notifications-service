package kafka

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

func PublishEvent(topic, key string, byteData []byte, host string) error {
	p := GetProducer(host)

	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          byteData,
	}, nil)

	if err != nil {
		log.Printf("Kafka publish failed: %v", err)
		return err
	}

	p.Flush(500)

	return nil
}
