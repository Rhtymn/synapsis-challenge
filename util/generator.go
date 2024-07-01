package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func CreateInvoice() string {
	t := time.Now()
	sb := strings.Builder{}
	sb.WriteString("INV/")
	sb.WriteString(GenerateRandomNumericString(4))
	sb.WriteString(fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day()))
	return sb.String()
}

func GenerateRandomNumericString(length int) string {
	const digits = "0123456789"
	rand.New(rand.NewSource(time.Now().UnixNano()))

	randomBytes := make([]byte, length)

	for i := 0; i < len(randomBytes); i++ {
		randomBytes[i] = digits[rand.Intn(len(digits))]
	}

	return string(randomBytes)
}
