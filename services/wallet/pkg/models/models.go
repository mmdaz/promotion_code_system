package models

import "time"

type Wallet struct {
	PhoneNumber string
	Cache       int
	UpdatedAt   time.Time
}
