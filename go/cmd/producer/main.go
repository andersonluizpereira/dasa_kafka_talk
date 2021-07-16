package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("transferiu", "teste", producer, []byte("transferencia"), deliveryChan)
	go DeliveryReport(deliveryChan)
	producer.Flush(10000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "localhost:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true",
	}

	//Em Acks=all, o produtor aguarda o retorno até que o líder e todas as réplicas recebam a mensagem.
	//É o modo mais seguro, mas lembre-se que para funcionar o tópico precisa ter um fator de replicação maior
	//que 1 e o cluster deve ter mais de 1 broker.

	//Em Acks=1, o produtor aguarda por um ok do líder da partição.
	//Sendo assim sabemos que ao menos 1 broker recebeu a mensagem.

	//Em Acks=0, o produtor não aguarda nenhum tipo de resposta do cluster.
	//É o modo com throughput mais elevado. É importante levar em conta que nesse modo a perda de dados é possível uma vez que o produtor não aguarda nenhum tipo de sinal do cluster.

	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			} else {
				fmt.Println("Mensagem enviada ", ev.TopicPartition)
				// anotar no banco de dados que a mensagem foi processado.
				// ex: confirma que uma transferencia bancaria ocorreu.
			}
		}
	}
}
