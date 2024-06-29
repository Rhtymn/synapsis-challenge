package util

import (
	"time"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	AccountID  int64 `json:"accountId"`
	Permission int64 `json:"permission"`
}

type JWTProvider interface {
	CreateToken(accountID int64) (string, error)
	VerifyToken(token string) (JWTClaims, error)
}

type jwtProviderHS256 struct {
	permission int64
	issuer     string
	secretKey  string
	lifespan   time.Duration
}

func NewJWTProvider(permission int64, issuer string, secretKey string, lifespan time.Duration) *jwtProviderHS256 {
	return &jwtProviderHS256{
		permission: permission,
		issuer:     issuer,
		secretKey:  secretKey,
		lifespan:   lifespan,
	}
}

func (p *jwtProviderHS256) CreateToken(accountID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    p.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(p.lifespan)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		AccountID:  accountID,
		Permission: p.permission,
	})

	signed, err := token.SignedString([]byte(p.secretKey))
	if err != nil {
		return "", apperror.Wrap(err)
	}

	return signed, nil
}

func (p *jwtProviderHS256) VerifyToken(tokenStr string) (JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&JWTClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(p.secretKey), nil
		},
		jwt.WithIssuer(p.issuer),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return JWTClaims{}, apperror.NewInvalidToken(err)
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return JWTClaims{}, apperror.NewTypeAssertionFailed(claims, token)
	}
	return *claims, nil
}

type jwtProviderAny struct {
	providers []JWTProvider
}

func NewJWTProviderAny(providers []JWTProvider) *jwtProviderAny {
	return &jwtProviderAny{
		providers: providers,
	}
}

func (p *jwtProviderAny) CreateToken(accountID int64) (string, error) {
	return "", apperror.NewInternalFmt("uninplemented")
}

func (p *jwtProviderAny) VerifyToken(tokenstr string) (JWTClaims, error) {
	var claims JWTClaims
	var err error

	for _, prov := range p.providers {
		claims, err = prov.VerifyToken(tokenstr)
		if err == nil {
			return claims, nil
		}
	}

	return JWTClaims{}, err
}
