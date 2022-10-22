package metamask

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/buley/n3ut/domain"
)

type Claims struct {
	AddressHex string `json:"address"`
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
