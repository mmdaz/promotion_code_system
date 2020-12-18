package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/gommon/log"

	"math/rand"
	"time"
)

type PubSub interface {
	Publish(topic string, value []byte) error
	Subscribe(topics []string) error
	ReadMessage(timeout time.Duration) (*Message, error)
}

type Message struct {
	Value     []byte
	Timestamp time.Time
	Topic     kafka.TopicPartition
}

type KafkaPubSub struct {
	consumer *kafka.Consumer
	producer *kafka.Producer
}

type KafkaOption struct {
	Servers           string
	GroupID           string
	OffsetReset       string
	DisableAutoCommit bool
}

func NewKafka(option KafkaOption) PubSub {
	rand.Seed(time.Now().Unix())
	var gID = fmt.Sprintf("groupid_%d", rand.Int31())
	if option.GroupID != "random" {
		gID = option.GroupID
	}
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  option.Servers,
		"group.id":           gID,
		"auto.offset.reset":  option.OffsetReset,
		"enable.auto.commit": !option.DisableAutoCommit,
	})
	if err != nil {
		log.Fatal(err)
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": option.Servers,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Delivery report handler for produced messages
	go func(p *kafka.Producer) {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Errorf("Delivery failed: %v\n", ev.TopicPartition)
				}
			}
		}
	}(producer)

	return &KafkaPubSub{
		consumer: consumer,
		producer: producer,
	}
}

func (pb *KafkaPubSub) Publish(topic string, value []byte) error {
	return pb.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)
}

func (pb *KafkaPubSub) Subscribe(topics []string) error {
	log.Debug("Subscribe to kafka topic:", topics)
	return pb.consumer.SubscribeTopics(topics, nil)
}

func (pb *KafkaPubSub) ReadMessage(timeout time.Duration) (*Message, error) {
	message, err := pb.consumer.ReadMessage(timeout)
	if err != nil {
		return nil, err
	}

	return &Message{Value: message.Value, Timestamp: message.Timestamp, Topic: message.TopicPartition}, nil
}
