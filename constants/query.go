package constants

const (
	UserColumns = ` id, name, photo_url, date_of_birth, gender, phone_number, id_account, main_address_id `

	AccountWithCredentialColumns = ` id, email, email_verified, password, account_type, profile_set `
	AccountColumns               = ` id, email, email_verified, account_type, profile_set `

	EmailVerifyTokenColumns = ` id, token, expired_at `

	UserAddressColumns = ` id, name, phone_number, address, latitude, longitude, id_user `
)
