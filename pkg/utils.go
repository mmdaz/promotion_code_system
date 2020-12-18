package pkg

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"github.com/pkg/errors"
	"strings"
)

type PhoneNumberEnvelop struct {
	Region      string
	CountryCode int32
	Number      int
}

func NormalizePhone(phone string, defaultCountry string) (*PhoneNumberEnvelop, error) {
	normalPhone := phone
	if !strings.HasPrefix(phone, "+") {
		normalPhone = fmt.Sprintf("+%v", phone)
	}
	num, err := phonenumbers.Parse(normalPhone, defaultCountry)
	if err != nil {
		return nil, err
	}

	if num.CountryCode == nil || num.NationalNumber == nil {
		return nil, errors.New("invalid phone number")
	}

	return &PhoneNumberEnvelop{
		CountryCode: *num.CountryCode,
		Number:      int(*num.NationalNumber),
		Region:      phonenumbers.GetRegionCodeForNumber(num),
	}, nil
}
