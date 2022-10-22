package users

import (
	"github.com/buley/n3ut/application/user"
	"github.com/buley/n3ut/interfaces/server"
)

func SetUp(cntl *server.Controller) {
	verifier := NewVerifier()
	cntl.GET("/:address", GetHandler, verifier)
	cntl.PUT("/:address", UpdateHandler, verifier)
	cntl.DELETE("/:address", DeleteHandler, verifier)
}

func GetHandler(c *server.Context) error {
	addressHex := c.Param("address")

	ctx := c.Request().Context()
	in := user.NewGetUserInput(addressHex)

	out, err := c.Apps.User.GetUser(ctx, in)
	if err != nil {
		return c.JSONError(err)
	}

	return c.JSONSuccess(out)
}

func UpdateHandler(c *server.Context) error {

	addressHex := c.FormValue("user_address")
	firstName := c.FormValue("user_first_name")
	lastName := c.FormValue("user_last_name")
	messageToAddress := c.FormValue("message_to_address")
	messageBody := c.FormValue("message_body")
	messageSubject := c.FormValue("message_subject")

	ctx := c.Request().Context()
	in := user.NewUpdateUserInput(addressHex, lastName, firstName, messageToAddress, messageBody, messageSubject)

	out, err := c.Apps.User.UpdateUser(ctx, in)
	if err != nil {
		return c.JSONError(err)
	}

	return c.JSONSuccess(out)
}

func DeleteHandler(c *server.Context) error {
	addressHex := c.Param("address")

	ctx := c.Request().Context()
	in := user.NewDeleteUserInput(addressHex)

	out, err := c.Apps.User.DeleteUser(ctx, in)
	if err != nil {
		return c.JSONError(err)
	}

	return c.JSONSuccess(out)
}
