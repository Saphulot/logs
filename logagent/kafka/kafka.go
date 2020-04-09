package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var producer sarama.SyncProducer

func Init(kafkaAddress []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	fmt.Println(kafkaAddress)
	producer, err = sarama.NewSyncProducer(kafkaAddress, config)
	if err != nil {
		fmt.Printf("init kafka fail : %v/n", err)
		return err
	}
	return
}

func SendMsg(topic string, data string) (err error) {

	msg := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(data)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Printf("FAILED to send message: %s\n", err)
		return err
	} else {
		fmt.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
	return
}
