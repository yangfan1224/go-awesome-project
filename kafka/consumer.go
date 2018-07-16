package kafka
import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consumer(){
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "servicenode01:9093,servicenode02:9093,servicenode03:9093,servicenode04:9093,servicenode05:9093",
		"group.id":          "myGroup",
		"auto.offset.reset": "latest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"userTagStatics"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()
}
