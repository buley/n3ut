package server

import (
	"github.com/buley/n3ut/application/auth"
	"github.com/buley/n3ut/application/user"
)

type Apps struct {
	Auth auth.Application
	User user.Application
}
