package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// 配置生产者
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// 创建生产者
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalln("Failed to start producer:", err)
	}
	defer producer.Close()

	// 创建消息
	msg := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	// 发送消息
	for i := 0; i < 10; i++ {
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatalln("Failed to send message:", err)
		}
		fmt.Println(i)
		log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
		time.Sleep(2 * time.Second)
	}
}