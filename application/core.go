package application

import (
	"github.com/buley/n3ut/domain/auth"
	"github.com/buley/n3ut/domain/user"
)

type Core struct {
	Config       *Config
	Services     *Services
	Repositories *Repositories
}

type Services struct {
	Auth auth.Service
}

type Repositories struct {
	User user.Repository
}
