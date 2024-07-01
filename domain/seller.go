package domain

import "time"

type Seller struct {
	ID          int64
	Name        string
	DateOfBirth *time.Time
	Gender      *string
	PhoneNumber *string
	Account     Account
}
