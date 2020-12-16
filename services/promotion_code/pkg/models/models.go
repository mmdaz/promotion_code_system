package models

import "time"

type PromotionCode struct {
	Value       string
	PhoneNumber string
	CreatedAt   time.Time
}
