package service

import (
	"context"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/constants"
	"github.com/Rhtymn/synapsis-challenge/domain"
	"github.com/Rhtymn/synapsis-challenge/util"
)

type AccountServiceOpts struct {
	Account              domain.AccountRepository
	User                 domain.UserRepository
	EmailVerifyToken     domain.EmailVerifyTokenRepository
	PasswordHasher       util.PasswordHasher
	Transactor           util.Transactor
	UserAccessProvider   util.JWTProvider
	SellerAccessProvider util.JWTProvider
	AdminAccessProvider  util.JWTProvider
}

type accountService struct {
	accountRepository          domain.AccountRepository
	userRepository             domain.UserRepository
	emailVerifyTokenRepository domain.EmailVerifyTokenRepository
	passwordHasher             util.PasswordHasher
	transactor                 util.Transactor
	userAccessProvider         util.JWTProvider
	sellerAccessProvider       util.JWTProvider
	adminAccessProvider        util.JWTProvider
}

func NewAccountService(opts AccountServiceOpts) *accountService {
	return &accountService{
		accountRepository:          opts.Account,
		userRepository:             opts.User,
		emailVerifyTokenRepository: opts.EmailVerifyToken,
		passwordHasher:             opts.PasswordHasher,
		transactor:                 opts.Transactor,
		userAccessProvider:         opts.UserAccessProvider,
		sellerAccessProvider:       opts.SellerAccessProvider,
		adminAccessProvider:        opts.AdminAccessProvider,
	}
}

func (s *accountService) Login(ctx context.Context, cred domain.AccountLoginCredentials) (domain.AuthToken, error) {
	a, err := s.accountRepository.GetByEmail(ctx, cred.Account.Email)
	if err != nil {
		if apperror.IsErrorCode(err, apperror.CodeNotFound) {
			return domain.AuthToken{}, apperror.NewUnauthorized(err, "invalid email or password")
		}
		return domain.AuthToken{}, nil
	}

	err = s.passwordHasher.CheckPassword(a.Password, cred.Password)
	if err != nil {
		if apperror.IsErrorCode(err, apperror.CodeWrongPassword) {
			return domain.AuthToken{}, apperror.NewUnauthorized(err, "invalid email or password")
		}
		return domain.AuthToken{}, nil
	}

	token, err := s.CreateTokensForAccount(a.Account.ID, a.Account.AccountType)
	if err != nil {
		return domain.AuthToken{}, apperror.Wrap(err)
	}
	return token, nil
}

func (s *accountService) Register(ctx context.Context, cred domain.AccountRegisterCredentials) (domain.Account, error) {
	var account domain.Account
	err := s.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		exist, err := s.accountRepository.IsEmailUsed(ctx, cred.Account.Email)
		if exist {
			return apperror.NewAlreadyExists(err, "email already used")
		}

		hashedPassword, err := s.passwordHasher.HashPassword(cred.Password)
		if err != nil {
			return err
		}
		cred.Password = hashedPassword

		account, err = s.accountRepository.Add(ctx, domain.AccountWithCredentials{
			Account:  cred.Account,
			Password: cred.Password,
		})
		if err != nil {
			return err
		}

		_, err = s.userRepository.Add(ctx, domain.User{
			Account: domain.Account{
				ID:          account.ID,
				AccountType: cred.Account.AccountType,
			},
			Name: util.EmailToName(cred.Account.Email),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return account, err
	}
	return account, nil
}

func (s *accountService) CreateTokensForAccount(accountID int64, accountType string) (domain.AuthToken, error) {
	var accessProvider util.JWTProvider

	if accountType == constants.USER {
		accessProvider = s.userAccessProvider
	} else if accountType == constants.SELLER {
		accessProvider = s.sellerAccessProvider
	} else {
		accessProvider = s.adminAccessProvider
	}

	accessToken, err := accessProvider.CreateToken(accountID)
	if err != nil {
		return domain.AuthToken{}, apperror.Wrap(err)
	}
	accessClaims, err := accessProvider.VerifyToken(accessToken)
	if err != nil {
		return domain.AuthToken{}, apperror.Wrap(err)
	}

	token := domain.AuthToken{
		AccessToken:     accessToken,
		AccessExpiredAt: accessClaims.ExpiresAt.Time,
	}
	return token, nil
}
