package mocks

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/mmdaz/arvan-challenge/pkg/kafka"
	"time"
)

type MockKafka struct {
	mock.Mock
	PublishFunc func(topic string, value []byte) error
}

func (m *MockKafka) Publish(topic string, value []byte) error {
	return m.PublishFunc(topic, value)
}

func (m *MockKafka) Subscribe(topics []string) error {
	panic("implement me")
}

func (m *MockKafka) ReadMessage(timeout time.Duration) (*kafka.Message, error) {
	panic("implement me")
}

func NewKafkaMock() *MockKafka {
	return &MockKafka{}
}
