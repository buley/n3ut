package metamask

import (
	"time"

	"github.com/buley/n3ut/domain"
	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	AddressHex string `json:"user_address"`
	jwt.StandardClaims
}

func newClaims(address domain.Address, d time.Duration) *Claims {
	now := time.Now()

	return &Claims{
		AddressHex: address.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(d).Unix(),
			IssuedAt:  now.Unix(),
		},
	}
}
