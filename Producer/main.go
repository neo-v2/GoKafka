package main

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka1:39091"})
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	//Produce message
	topic := "Test1" // Or the topic name you created
	for _, v := range []string{
		"Hello World , Welcome to My Kafka",
	} {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(v),
		}, nil)
	}

}
