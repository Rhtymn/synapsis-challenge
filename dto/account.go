package dto

import (
	"strings"

	"github.com/Rhtymn/synapsis-challenge/domain"
)

type AccountRegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=24"`
}

type AccountLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=24"`
}

type AccountTypeParams struct {
	Type string `uri:"type" binding:"required,oneof=user seller"`
}

func (ar *AccountRegisterRequest) ToCredentials(accountType string) domain.AccountRegisterCredentials {
	return domain.AccountRegisterCredentials{
		Account: domain.Account{
			Email:       strings.ToLower(ar.Email),
			AccountType: accountType,
		},
		Password: ar.Password,
	}
}

func (al *AccountLoginRequest) ToCredentials() domain.AccountLoginCredentials {
	return domain.AccountLoginCredentials{
		Account: domain.Account{
			Email: strings.ToLower(al.Email),
		},
		Password: al.Password,
	}
}
