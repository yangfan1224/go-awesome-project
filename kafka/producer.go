package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"time"
)

func ProducerExample(){
	broker := "dmp-test04:9092,dmp-test05:9092,dmp-test06:9092"
	topic := "UrlTagMessage"

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v\n", p)

	doneChan := make(chan bool)

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v, key %s, value %s\n",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset, string(m.Key), string(m.Value))
				}
			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()
	//Url 	string `json:"url"`
	//Host 	string `json:host`
	//TagOne 	string `json:"cate1"`
	//TagTwo  string `json:"cate2"`
	//Timestamps int64 `json:"timestamps"`
	jsonMessage := `{"url":"https://blog.csdn.net/yuyinghua0302/article/details/78612668", "host":"blog.csdn.net", "cate1":"cate1_17", "cate2":"cate2_49", "timestamps": 1543402764333}`
	for i := 0; i< 1; i++{
		p.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(jsonMessage)}
		time.Sleep(time.Millisecond)
	}


	// wait for delivery report goroutine to finish
	_ = <-doneChan

	p.Close()
}
