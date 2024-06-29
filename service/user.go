package service

import (
	"context"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/domain"
	"github.com/Rhtymn/synapsis-challenge/util"
)

type userService struct {
	userRepository        domain.UserRepository
	userAddressRepository domain.UserAddressRepository
	transactor            util.Transactor
}

type UserServiceOpts struct {
	User        domain.UserRepository
	UserAddress domain.UserAddressRepository
	Transactor  util.Transactor
}

func NewUserService(opts UserServiceOpts) *userService {
	return &userService{
		userRepository:        opts.User,
		userAddressRepository: opts.UserAddress,
		transactor:            opts.Transactor,
	}
}

func (s *userService) AddAddress(ctx context.Context, ua domain.UserAddress) (domain.UserAddress, error) {
	var userAddress domain.UserAddress
	err := s.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		accountID, err := util.GetAccountIDFromContext(ctx)
		if err != nil {
			return apperror.Wrap(err)
		}

		user, err := s.userRepository.GetByAccountID(ctx, accountID)
		if err != nil {
			return apperror.Wrap(err)
		}

		ua.User.ID = user.ID
		userAddress, err = s.userAddressRepository.Add(ctx, ua)
		if err != nil {
			return apperror.Wrap(err)
		}

		if user.MainAddressID == nil {
			err := s.userRepository.SetMainAddressByID(ctx, userAddress.ID, user.ID)
			if err != nil {
				return apperror.Wrap(err)
			}
		}
		return nil
	})
	if err != nil {
		return userAddress, apperror.Wrap(err)
	}
	return userAddress, nil
}

func (s *userService) UpdateMainAddress(ctx context.Context, addressID int64) error {
	accountID, err := util.GetAccountIDFromContext(ctx)
	if err != nil {
		return apperror.Wrap(err)
	}

	user, err := s.userRepository.GetByAccountID(ctx, accountID)
	if err != nil {
		return apperror.Wrap(err)
	}

	if user.MainAddressID != nil {
		m := *user.MainAddressID
		if m == addressID {
			return apperror.NewBadRequest(nil, "address already set to main address")
		}
	}

	userAddress, err := s.userAddressRepository.GetByID(ctx, addressID)
	if err != nil {
		return apperror.Wrap(err)
	}

	if userAddress.User.ID != user.ID {
		return apperror.NewNotFound(nil, "user address not found")
	}

	err = s.userRepository.SetMainAddressByID(ctx, addressID, user.ID)
	if err != nil {
		return apperror.Wrap(err)
	}
	return nil
}
