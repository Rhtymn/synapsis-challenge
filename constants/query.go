package constants

const (
	UserColumns = ` id, name, photo_url, date_of_birth, gender, phone_number, id_account `

	AccountWithCredentialColumns = ` id, email, email_verified, password, account_type, profile_set `
	AccountColumns               = ` id, email, email_verified, account_type, profile_set `

	EmailVerifyTokenColumns = ` id, token, expired_at `
)
