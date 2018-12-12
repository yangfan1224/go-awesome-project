package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consumer(){
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "servicenode01:9093,servicenode02:9093,servicenode03:9093,servicenode04:9093,servicenode05:9093",
		"group.id":          "dsctest01",
		"auto.offset.reset": "latest",
		"session.timeout.ms":   6000,
		"enable.auto.commit": false,
	})

	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{"userTagStatics"}, nil)
	if err != nil{
		fmt.Printf("Consumer error: %v \n", err)
		panic(err)
	}
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on Partition = %d, offset = %d: %s\n", msg.TopicPartition.Partition, msg.TopicPartition.Offset, string(msg.Value))
		} else {
			if kafkaErr := err.(kafka.Error); kafkaErr.Code() == -185{
				fmt.Printf("read tiemout, code is %d", kafkaErr.Code())
				break
			}
			fmt.Printf("Consumer error: %d (%v)\n", err, msg)
			break
		}
	}

	c.Close()
}
