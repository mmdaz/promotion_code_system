package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestApplyCode(t *testing.T) {
	testPhoneNumber := 989147569102
	testCode := "asdfgh"
	promotionCodeRepo.CreateFunc = func(phoneNumber int, code string) error {
		assert.Equal(t, phoneNumber, 9147569102)
		assert.Equal(t, code, testCode)
		return nil
	}

	promotionCodeRepo.GetAppliedCodeCountsFunc = func(startTime time.Time, endTime time.Time) (int32, error) {
		return 1, nil
	}

	kafkaMock.PublishFunc = func(topic string, value []byte) error {
		//assert.Equal(t, phoneNumber, testPhoneNumber)
		return nil
	}

	walletMock.IncreaseAmountFunc = func(phoneNumber int, amount int) error {
		assert.Equal(t, phoneNumber, 9147569102)
		assert.Equal(t, amount, 1000)
		return nil
	}

	err := promotionCodeCore.ApplyPromotionCode(fmt.Sprint(testPhoneNumber))
	assert.Nil(t, err)
}
