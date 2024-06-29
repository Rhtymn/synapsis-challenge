package domain

import (
	"context"
	"mime/multipart"
	"time"
)

type User struct {
	ID          int64
	Account     Account
	Name        string
	PhotoURL    *string
	DateOfBirth *time.Time
	Gender      *string
	PhoneNumber *string
}

type UserProfile struct {
	Name        string
	PhotoURL    multipart.File
	DateOfBirth time.Time
	Gender      string
	PhoneNumber string
}

type UserRepository interface {
	GetById(ctx context.Context, id int64) (User, error)
	GetByIdAndLock(ctx context.Context, id int64) (User, error)

	Add(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) (User, error)
	IsPhoneNumberUsed(ctx context.Context, phoneNumber string) (bool, error)
}

type UserService interface {
	CreateProfile(ctx context.Context, up UserProfile) (User, error)
}
