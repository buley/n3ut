package handler

import (
	"github.com/buley/n3ut/interfaces/server"
	"github.com/buley/n3ut/interfaces/server/handler/api"
	"github.com/buley/n3ut/interfaces/server/handler/auth"
)

func SetUp(cntl *server.Controller) {
	auth.SetUp(cntl.Child("/auth"))
	api.SetUp(cntl.Child("/api"))
}
