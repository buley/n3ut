package user

import "github.com/buley/n3ut/domain"

type GetUserOutput struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	AddressHex string `json:"address"`
}

func NewGetUserOutput(u *domain.User) *GetUserOutput {
	return &GetUserOutput{
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		AddressHex: u.Address.Hex(),
	}
}

type UpdateUserOutput struct{}

func NewUpdateUserOutput() *UpdateUserOutput {
	return &UpdateUserOutput{}
}

type DeleteUserOutput struct{}

func NewDeleteUserOutput() *DeleteUserOutput {
	return &DeleteUserOutput{}
}
