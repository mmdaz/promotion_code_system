package models

import "time"

type PromotionCode struct {
	Value       string
	PhoneNumber int
	CreatedAt   time.Time
}
