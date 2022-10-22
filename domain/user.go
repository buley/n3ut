package domain

import (
	"unicode/utf8"
)

type User struct {
	FirstName string
	LastName  string
	Address   Address
	Challenge string
}

type UserMessage struct {
	FirstName string
	LastName  string
	Address   Address
	Challenge string
}

func NewUser(address Address, lastName string, firstName string) *User {
	return &User{
		FirstName: firstName,
		Address:   address,
		LastName:  lastName,
	}
}

func ValidateUserFirstName(firstName string) error {
	l := utf8.RuneCountInString(firstName)
	if l < UserNameLengthMin {
		return ErrTooShortUserName
	}
	if l > UserNameLengthMax {
		return ErrTooLongUserName
	}
	return nil
}

func ValidateUserLastName(lastName string) error {
	l := utf8.RuneCountInString(lastName)
	if l < UserNameLengthMin {
		return ErrTooShortUserName
	}
	if l > UserNameLengthMax {
		return ErrTooLongUserName
	}
	return nil
}
