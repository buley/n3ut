package application

import "github.com/buley/n3ut/domain"

type AddressHexInput struct {
	AddressHex string
}

func NewAddressHexInput(addressHex string) *AddressHexInput {
	return &AddressHexInput{
		AddressHex: addressHex,
	}
}

func (in *AddressHexInput) Validate() error {
	if err := domain.ValidateAddressHex(in.AddressHex); err != nil {
		return err
	}
	return nil
}

func (in *AddressHexInput) Address() domain.Address {
	return domain.NewAddressFromHex(in.AddressHex)
}
