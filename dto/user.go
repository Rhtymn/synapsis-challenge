package dto

import "github.com/Rhtymn/synapsis-challenge/domain"

type CreateAddressRequest struct {
	Name        string        `json:"name" binding:"required"`
	PhoneNumber string        `json:"phone_number" binding:"required,min=10,max=15,number"`
	Address     string        `json:"address" binding:"required"`
	Coordinate  CoordinateDTO `json:"coordinate" binding:"required"`
}

type UserAddressResponse struct {
	Name        string        `json:"name" binding:"required"`
	PhoneNumber string        `json:"phone_number" binding:"required,min=10,max=15"`
	Address     string        `json:"address" binding:"required"`
	Coordinate  CoordinateDTO `json:"coordinate" binding:"required"`
}

type UpdateMainAddressParams struct {
	AddressID int64 `uri:"address_id" binding:"required"`
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
