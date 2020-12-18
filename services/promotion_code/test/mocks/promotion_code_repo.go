package mocks

import (
	"github.com/stretchr/testify/mock"
	"time"
)

type MockPromotionCodeRepo struct {
	mock.Mock
	CreateFunc               func(phoneNumber int, code string) error
	GetAppliedCodeCountsFunc func(startTime time.Time, endTime time.Time) (int32, error)
}

func NewMockPromotionCodeRepo() *MockPromotionCodeRepo {
	return &MockPromotionCodeRepo{}
}

func (mp *MockPromotionCodeRepo) Create(phoneNumber int, code string) error {
	return mp.CreateFunc(phoneNumber, code)

}
func (mp *MockPromotionCodeRepo) GetAppliedCodeCounts(startTime time.Time, endTime time.Time) (int32, error) {
	return mp.GetAppliedCodeCountsFunc(startTime, endTime)
}
