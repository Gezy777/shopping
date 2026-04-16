package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// 创建消费者
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "my-group", config)
	if err != nil {
		log.Fatalln("Failed to start consumer:", err)
	}
	defer consumer.Close()

	// 订阅主题
	partitionConsumer, err := consumer.ConsumePartition("test-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln("Failed to start partition consumer:", err)
	}
	defer partitionConsumer.Close()

	// 消费消息
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Received message: %s\n", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			log.Println("Error:", err)
		}
	}
}