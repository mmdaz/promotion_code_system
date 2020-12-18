package internal

import (
	"github.com/labstack/gommon/log"
	"gitlab.com/mmdaz/arvan-challenge/pkg"
	"gitlab.com/mmdaz/arvan-challenge/pkg/kafka"
	"strconv"
)

type Publisher struct {
	kafka      *kafka.PubSub
	config     *pkg.Config
	updateChan chan int
}

func NewPublisher(kafka *kafka.PubSub, config *pkg.Config) *Publisher {
	p := &Publisher{kafka: kafka, config: config, updateChan: make(chan int)}
	p.Run()
	return p
}

func (p Publisher) ReceiveUpdate(phoneNumber int) {
	p.updateChan <- phoneNumber
}

func (p *Publisher) Run() {
	for i := 0; i < 100; i++ {
		go p.worker()
	}
}

func (p *Publisher) worker() {
	for {
		phoneNumber := <-p.updateChan
		err := p.publish(phoneNumber)
		if err != nil {
			log.Error(err)
		}
	}
}

func (p *Publisher) publish(phoneNumber int) error {
	log.Info("publish: ", phoneNumber, p.config.Kafka.Topic)
	bytes := []byte(strconv.Itoa(phoneNumber))
	err := p.kafka.Publish(p.config.Kafka.Topic, bytes)
	return err
}
