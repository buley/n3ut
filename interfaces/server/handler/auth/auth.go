package auth

import (
	"github.com/buley/n3ut/application/auth"
	"github.com/buley/n3ut/interfaces/server"
)

func SetUp(cntl *server.Controller) {
	cntl.POST("/challenge", ChallengeHandler)
	cntl.POST("/authorize", AuthorizeHandler)
}

func ChallengeHandler(c *server.Context) error {
	addressHex := c.FormValue("user_address")

	ctx := c.Request().Context()
	in := auth.NewChallengeInput(addressHex)

	out, err := c.Apps.Auth.Challenge(ctx, in)
	if err != nil {
		return c.JSONError(err)
	}

	return c.JSONSuccess(out)
}

func AuthorizeHandler(c *server.Context) error {
	addressHex := c.FormValue("user_address")
	sigHex := c.FormValue("user_signature")

	ctx := c.Request().Context()
	in := auth.NewAuthorizeInput(addressHex, sigHex)

	out, err := c.Apps.Auth.Authorize(ctx, in)
	if err != nil {
		return c.JSONError(err)
	}

	return c.JSONSuccess(out)
}
