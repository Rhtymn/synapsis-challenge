package util

func EmailToName(email string) string {
	idx := 0
	for i := 0; i < len(email); i++ {
		if email[i] == '@' {
			idx = i
		}
	}
	return email[:idx]
}
