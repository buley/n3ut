package api

import (
	"github.com/buley/n3ut/infrastructure/auth/metamask"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewAuthenticator(secret string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &metamask.Claims{},
		SigningKey: []byte(secret),
	})
}
