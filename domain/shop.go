package domain

type Shop struct {
	ID          int64
	ShopName    string
	Slug        string
	LogoURL     *string
	PhoneNumber string
	Description *string
	Address     string
	Latitude    float64
	Longitude   float64
	IsActive    bool
	Seller      Seller
}
