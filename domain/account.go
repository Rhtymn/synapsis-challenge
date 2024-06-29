package domain

import "context"

const (
	AccountTypeUser   = "user"
	AccountTypeSeller = "seller"
)

type Account struct {
	ID            int64
	Email         string
	EmailVerified bool
	AccountType   string
	ProfileSet    bool
}

type AccountWithCredentials struct {
	Account  Account
	Password string
}

type AccountLoginCredentials struct {
	Account  Account
	Password string
}

type AccountRegisterCredentials struct {
	Account  Account
	Password string
}

type AccountVerifyEmailCredentials struct {
	EmailVerifyToken string
}

type AccountRepository interface {
	GetById(ctx context.Context, id int64) (AccountWithCredentials, error)
	GetByIdAndLock(ctx context.Context, id int64) (AccountWithCredentials, error)
	GetByEmail(ctx context.Context, email string) (AccountWithCredentials, error)
	GetByEmailAndLock(ctx context.Context, email string) (AccountWithCredentials, error)

	Add(ctx context.Context, a AccountWithCredentials) (Account, error)
	VerifyEmailById(ctx context.Context, id int64) error
	ProfileSetById(ctx context.Context, id int64) error

	IsEmailUsed(ctx context.Context, email string) (bool, error)
}

type AccountService interface {
	Login(ctx context.Context, cred AccountLoginCredentials) (AuthToken, error)
	Register(ctx context.Context, cred AccountRegisterCredentials) (Account, error)

	CreateTokensForAccount(accountID int64, accountType string) (AuthToken, error)

	GetVerifyEmailToken(ctx context.Context) error
	CheckVerifyEmailToken(ctx context.Context, email string, token string) error
	VerifyEmail(ctx context.Context, email string, token string) error
}
