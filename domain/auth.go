package domain

import "time"

type AuthToken struct {
	AccessToken     string
	AccessExpiredAt time.Time
}
