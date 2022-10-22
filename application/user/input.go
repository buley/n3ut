package user

import (
	"github.com/buley/n3ut/application"
	"github.com/buley/n3ut/domain"
)

type GetUserInput struct {
	*application.AddressHexInput
}

func NewGetUserInput(addressHex string) *GetUserInput {
	return &GetUserInput{
		AddressHexInput: application.NewAddressHexInput(addressHex),
	}
}

type UpdateUserInput struct {
	*application.AddressHexInput
	FirstName string
	LastName  string
	Message   struct {
		ToAddress string
		Subject   string
		Body      string
	}
}

func NewUpdateUserInput(addressHex, lastName string, firstName string, toAddress string, subject string, body string) *UpdateUserInput {
	return &UpdateUserInput{
		AddressHexInput: application.NewAddressHexInput(addressHex),
		FirstName:       firstName,
		LastName:        lastName,
		Message: struct {
			ToAddress string
			Subject   string
			Body      string
		}(NeutralMessage{
			ToAddress: toAddress,
			Subject:   subject,
			Body:      body,
		}),
	}
}

func (in *UpdateUserInput) Validate() error {
	if err := in.AddressHexInput.Validate(); err != nil {
		return err
	}
	if err := domain.ValidateUserFirstName(in.FirstName); err != nil {
		return err
	}
	if err := domain.ValidateUserLastName(in.LastName); err != nil {
		return err
	}
	return nil
}

type DeleteUserInput struct {
	*application.AddressHexInput
}

func NewDeleteUserInput(addressHex string) *DeleteUserInput {
	return &DeleteUserInput{
		AddressHexInput: application.NewAddressHexInput(addressHex),
	}
}
