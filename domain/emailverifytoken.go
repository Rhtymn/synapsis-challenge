package domain

import (
	"context"
	"time"
)

type EmailVerifyToken struct {
	ID        int64
	Token     string
	ExpiredAt time.Time
	Account   Account
}

type EmailVerifyTokenRepository interface {
	GetById(ctx context.Context, id int64) (EmailVerifyToken, error)
	GetByTokenStr(ctx context.Context, tokenStr string) (EmailVerifyToken, error)
	GetByTokenStrAndLock(ctx context.Context, tokenStr string) (EmailVerifyToken, error)
	Add(ctx context.Context, e EmailVerifyToken) (EmailVerifyToken, error)
	SoftDeleteByToken(ctx context.Context, token string) error
}
