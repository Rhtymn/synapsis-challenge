package dto

import (
	"mime/multipart"

	"github.com/Rhtymn/synapsis-challenge/domain"
)

type CreateAddressRequest struct {
	Name        string        `json:"name" binding:"required"`
	PhoneNumber string        `json:"phone_number" binding:"required,min=10,max=15,number"`
	Address     string        `json:"address" binding:"required"`
	Coordinate  CoordinateDTO `json:"coordinate" binding:"required"`
}

type UserAddressResponse struct {
	Name        string        `json:"name"`
	PhoneNumber string        `json:"phone_number"`
	Address     string        `json:"address"`
	Coordinate  CoordinateDTO `json:"coordinate"`
}

type UpdateMainAddressParams struct {
	AddressID int64 `uri:"address_id" binding:"required"`
}

type UpdateProfileRequest struct {
	Name        string                `form:"name" binding:"required"`
	DateOfBirth string                `form:"date_of_birth" binding:"required"`
	Photo       *multipart.FileHeader `form:"photo"`
	PhoneNumber string                `form:"phone_number" binding:"required,min=10,max=15,number"`
	Gender      string                `form:"gender" binding:"required,oneof=male female"`
}

type UserResponse struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	PhotoURL      *string `json:"photo_url,omitempty"`
	DateOfBirth   *string `json:"date_of_birth,omitempty"`
	Gender        *string `json:"gender,omitempty"`
	PhoneNumber   *string `json:"phone_number,omitempty"`
	MainAddressID *int64  `json:"main_address_id,omitempty"`
}

func (c *CreateAddressRequest) ToUserAddressDomain() domain.UserAddress {
	return domain.UserAddress{
		Name:        c.Name,
		PhoneNumber: c.PhoneNumber,
		Address:     c.Address,
		Coordinate:  c.Coordinate.ToCoordinate(),
	}
}

func NewUserAddressResponse(ua domain.UserAddress) UserAddressResponse {
	return UserAddressResponse{
		Name:        ua.Name,
		PhoneNumber: ua.PhoneNumber,
		Address:     ua.Address,
		Coordinate:  NewCoordinateDTO(ua.Coordinate),
	}
}

func NewUserResponse(u domain.User) UserResponse {
	var dob *string
	if u.DateOfBirth != nil {
		t := u.DateOfBirth.Format("2006-01-02")
		dob = &t
	}
	return UserResponse{
		ID:            u.ID,
		Name:          u.Name,
		PhotoURL:      u.PhotoURL,
		DateOfBirth:   dob,
		Gender:        u.Gender,
		PhoneNumber:   u.PhoneNumber,
		MainAddressID: u.MainAddressID,
	}
}
