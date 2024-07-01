package service

import (
	"context"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/domain"
	"github.com/Rhtymn/synapsis-challenge/util"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type userService struct {
	userRepository        domain.UserRepository
	accountRepository     domain.AccountRepository
	userAddressRepository domain.UserAddressRepository
	transactor            util.Transactor
	cloudinaryProvider    util.CloudinaryProvider
}

type UserServiceOpts struct {
	User               domain.UserRepository
	UserAddress        domain.UserAddressRepository
	Account            domain.AccountRepository
	Transactor         util.Transactor
	CloudinaryProvider util.CloudinaryProvider
}

func NewUserService(opts UserServiceOpts) *userService {
	return &userService{
		userRepository:        opts.User,
		userAddressRepository: opts.UserAddress,
		accountRepository:     opts.Account,
		transactor:            opts.Transactor,
		cloudinaryProvider:    opts.CloudinaryProvider,
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

func (s *userService) UpdateProfile(ctx context.Context, up domain.UserProfile) (domain.User, error) {
	var user domain.User
	err := s.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		accountID, err := util.GetAccountIDFromContext(ctx)
		if err != nil {
			return apperror.Wrap(err)
		}

		a, err := s.accountRepository.GetById(ctx, accountID)
		if err != nil {
			return apperror.Wrap(err)
		}

		u, err := s.userRepository.GetByAccountID(ctx, accountID)
		if err != nil {
			return apperror.Wrap(err)
		}

		var photoURL *string
		if up.Photo != nil {
			res, err := s.cloudinaryProvider.Upload(ctx, *up.Photo, uploader.UploadParams{})
			if err == nil {
				photoURL = &res.SecureURL
			}
		}

		updatedUser, err := s.userRepository.Update(ctx, domain.User{
			ID:          u.ID,
			Name:        up.Name,
			DateOfBirth: &up.DateOfBirth,
			PhotoURL:    photoURL,
			Gender:      &up.Gender,
			PhoneNumber: &up.PhoneNumber,
		})
		if err != nil {
			return apperror.Wrap(err)
		}

		if !a.Account.ProfileSet {
			err := s.accountRepository.ProfileSetById(ctx, accountID)
			if err != nil {
				return apperror.Wrap(err)
			}
		}

		user = updatedUser
		return nil
	})
	if err != nil {
		return user, apperror.Wrap(err)
	}
	return user, nil
}
