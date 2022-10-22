package interfaces

import (
	"github.com/buley/n3ut/application"
	appAuth "github.com/buley/n3ut/application/auth"
	appUser "github.com/buley/n3ut/application/user"
	"github.com/buley/n3ut/infrastructure/auth/metamask"
	cacheUser "github.com/buley/n3ut/infrastructure/cache/user"
	"github.com/buley/n3ut/interfaces/server"
	"github.com/buley/n3ut/interfaces/server/handler"
)

func NewServer(conf *server.Config) *server.Server {
	srv := server.New(newCore(conf))

	srv.File("/", srv.Config.IndexFilePath)
	srv.Static("/static", srv.Config.StaticDirPath)

	handler.SetUp(srv.Base())

	return srv
}

func newCore(conf *server.Config) *server.Core {
	appCore := newAppCore(conf.App)

	return &server.Core{
		Config: conf,
		Apps: &server.Apps{
			Auth: appAuth.NewApplication(appCore),
			User: appUser.NewApplication(appCore),
		},
	}
}

func newAppCore(conf *application.Config) *application.Core {
	return &application.Core{
		Services: &application.Services{
			Auth: metamask.NewService(
				conf.Auth.Secret,
				conf.Auth.TokenExpiryDuration(),
			),
		},
		Repositories: &application.Repositories{
			User: cacheUser.NewRepository(),
		},
	}
}
