package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "dasa-consumer",
		"group.id":          "dasa-group",
		"auto.offset.reset": "earliest",
		//a mais antiga
		//latest - a mais recente
	}
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("erro Consumer", err.Error())
	}
	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
