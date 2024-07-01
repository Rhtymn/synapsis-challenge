package dto

import (
	"time"

	"github.com/Rhtymn/synapsis-challenge/domain"
)

type AuthTokenResponse struct {
	AccessToken     string    `json:"access_token"`
	AccessExpiredAt time.Time `json:"access_expired_at"`
}

func NewAuthTokenResponse(t domain.AuthToken) Response {
	return ResponseOK(AuthTokenResponse{
		AccessToken:     t.AccessToken,
		AccessExpiredAt: t.AccessExpiredAt,
	})
}
