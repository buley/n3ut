package auth

import "github.com/buley/n3ut/domain"

type Service interface {
	SetUpChallenge(u *domain.User) error
	VerifyResponse(u *domain.User, responseBytes []byte) error
	IssueToken(u *domain.User) ([]byte, error)
}
