package user

import "github.com/buley/n3ut/domain"

type GetUserOutput struct {
	FirstName  string `json:"userFirstName"`
	LastName   string `json:"userLastName"`
	AddressHex string `json:"userAddress"`
}

type NeutralMessage struct {
	ToAddress string `json:"messageToAddress"`
	Subject   string `json:"messageSubject"`
	Body      string `json:"messageBody"`
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
