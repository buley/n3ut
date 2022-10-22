package api

import (
	"github.com/buley/n3ut/interfaces/server"
	"github.com/buley/n3ut/interfaces/server/handler/api/users"
)

func SetUp(cntl *server.Controller) {
	authenticator := NewAuthenticator(cntl.Config.App.Auth.Secret)
	cntl.Use(authenticator)

	users.SetUp(cntl.Child("/users"))
}
