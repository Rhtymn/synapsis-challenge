package constants

const (
	UserColumns = ` id, name, photo_url, date_of_birth, gender, phone_number, id_account, main_address_id `

	AccountWithCredentialColumns = ` id, email, email_verified, password, account_type, profile_set `
	AccountColumns               = ` id, email, email_verified, account_type, profile_set `

	EmailVerifyTokenColumns = ` id, token, expired_at `

	UserAddressColumns = ` id, name, phone_number, address, latitude, longitude, id_user `

	ProductJoinedShopColumns = ` p.id, p.name, p.slug, p.photo_url, p.price, p.description, p.stock, p.id_shop, s.shop_name, s.slug `

	PaymentMethodColumns  = ` id, name `
	ShipmentMethodColumns = ` id, name `

	TransactionColumns = ` id, invoice, status, num_of_items, subtotal, 
							shipment_fee, total_fee, address, 
							latitude, longitude, phone_number,
							id_shop, id_user, id_shipment_method, id_payment_method `
	TransactionItemColumns = ` id, amount, total_price, id_transaction, id_product `

	PaymentColumns = ` id, file_url, id_transaction `

	SortASC  = "ASC"
	SortDESC = "DESC"
)
