package dto

import "github.com/Rhtymn/synapsis-challenge/domain"

type ShopDTO struct {
	ID          int64   `json:"id,omitempty"`
	ShopName    string  `json:"shop_name,omitempty"`
	Slug        string  `json:"slug,omitempty"`
	LogoURL     *string `json:"logo_url,omitempty"`
	PhoneNumber string  `json:"phone_number,omitempty"`
	Description *string `json:"description,omitempty"`
	Address     string  `json:"address,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
	IsActive    bool    `json:"is_active,omitempty"`
}

func NewShopDTO(s domain.Shop) ShopDTO {
	return ShopDTO{
		ID:          s.ID,
		ShopName:    s.ShopName,
		Slug:        s.Slug,
		LogoURL:     s.LogoURL,
		PhoneNumber: s.PhoneNumber,
		Description: s.Description,
		Address:     s.Address,
		Latitude:    s.Latitude,
		Longitude:   s.Longitude,
		IsActive:    s.IsActive,
	}
}
