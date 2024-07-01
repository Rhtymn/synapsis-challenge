package service

import (
	"context"
	"time"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/constants"
	"github.com/Rhtymn/synapsis-challenge/domain"
	"github.com/Rhtymn/synapsis-challenge/util"
)

type AccountServiceOpts struct {
	Account          domain.AccountRepository
	User             domain.UserRepository
	EmailVerifyToken domain.EmailVerifyTokenRepository
	PasswordHasher   util.PasswordHasher
	Transactor       util.Transactor

	UserAccessProvider   util.JWTProvider
	SellerAccessProvider util.JWTProvider
	AdminAccessProvider  util.JWTProvider
	RandomTokenProvider  util.RandomTokenProvider

	AppEmail      util.AppEmail
	EmailProvider util.EmailProvider
}

type accountService struct {
	accountRepository          domain.AccountRepository
	userRepository             domain.UserRepository
	emailVerifyTokenRepository domain.EmailVerifyTokenRepository

	passwordHasher util.PasswordHasher
	transactor     util.Transactor

	userAccessProvider   util.JWTProvider
	sellerAccessProvider util.JWTProvider
	adminAccessProvider  util.JWTProvider
	randomTokenProvider  util.RandomTokenProvider

	appEmail      util.AppEmail
	emailProvider util.EmailProvider
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
		randomTokenProvider:        opts.RandomTokenProvider,
		emailProvider:              opts.EmailProvider,
		appEmail:                   opts.AppEmail,
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

func (s *accountService) GetVerifyEmailToken(ctx context.Context) error {
	err := s.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		accountId, err := util.GetAccountIDFromContext(ctx)
		if err != nil {
			return apperror.Wrap(err)
		}

		err = s.emailVerifyTokenRepository.SoftDeleteByAccountID(ctx, accountId)
		if err != nil {
			return apperror.Wrap(err)
		}

		u, err := s.userRepository.GetByAccountID(ctx, accountId)
		if err != nil {
			return apperror.Wrap(err)
		}

		a, err := s.accountRepository.GetById(ctx, accountId)
		if err != nil {
			return apperror.Wrap(err)
		}

		if a.Account.EmailVerified {
			return apperror.NewAlreadyVerified("account already verified")
		}

		randomToken, err := s.randomTokenProvider.GenerateToken()
		if err != nil {
			return apperror.Wrap(err)
		}

		_, err = s.emailVerifyTokenRepository.Add(ctx, domain.EmailVerifyToken{
			Token:     randomToken,
			ExpiredAt: time.Now().Add(5 * time.Minute),
			Account: domain.Account{
				ID: a.Account.ID,
			},
		})
		if err != nil {
			return apperror.Wrap(err)
		}

		err = s.emailProvider.SendEmail(a.Account.Email, s.appEmail.NewVerifyAccountEmail(u.Name, a.Account.Email, randomToken))
		if err != nil {
			return apperror.Wrap(err)
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *accountService) CheckVerifyEmailToken(ctx context.Context, email string, token string) error {
	a, err := s.accountRepository.GetByEmail(ctx, email)
	if err != nil {
		return apperror.Wrap(err)
	}

	if a.Account.EmailVerified {
		return apperror.NewAlreadyVerified("account already verified")
	}

	emailVerifyToken, err := s.emailVerifyTokenRepository.GetByTokenStr(ctx, token)
	if err != nil {
		return apperror.NewInvalidVerifyEmailToken(err)
	}

	expiredAtLocal := util.ToLocalTime(emailVerifyToken.ExpiredAt)
	if expiredAtLocal.Before(time.Now()) {
		return apperror.NewBadRequest(nil, "token expired")
	}

	return nil
}

func (s *accountService) VerifyEmail(ctx context.Context, email string, token string) error {
	err := s.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		a, err := s.accountRepository.GetByEmail(ctx, email)
		if err != nil {
			return apperror.Wrap(err)
		}

		if a.Account.EmailVerified {
			return apperror.NewAlreadyVerified("account already verified")
		}

		emailVerifyToken, err := s.emailVerifyTokenRepository.GetByTokenStr(ctx, token)
		if err != nil {
			return apperror.NewInvalidVerifyEmailToken(err)
		}

		expiredAtLocal := util.ToLocalTime(emailVerifyToken.ExpiredAt)
		if expiredAtLocal.Before(time.Now()) {
			return apperror.NewBadRequest(nil, "token expired")
		}

		err = s.accountRepository.VerifyEmailById(ctx, a.Account.ID)
		if err != nil {
			return apperror.Wrap(err)
		}

		err = s.emailVerifyTokenRepository.SoftDeleteByToken(ctx, token)
		if err != nil {
			return apperror.Wrap(err)
		}

		return nil
	})
	if err != nil {
		return apperror.Wrap(err)
	}
	return nil
}
