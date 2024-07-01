package util

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/Rhtymn/synapsis-challenge/apperror"
)

type RandomTokenProvider interface {
	GenerateToken() (string, error)
}

type randomTokenImpl struct {
	length int
}

func NewRandomTokenProvider(length int) *randomTokenImpl {
	return &randomTokenImpl{
		length: length,
	}
}

func (p *randomTokenImpl) GenerateToken() (string, error) {
	b := make([]byte, p.length)
	_, err := rand.Read(b)
	if err != nil {
		return "", apperror.Wrap(err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
