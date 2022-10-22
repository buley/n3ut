package users

import (
	"github.com/buley/n3ut/interfaces/server"
	"github.com/labstack/echo"
)

func NewVerifier() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			c := &server.Context{
				Context: ec,
			}
			if c.Param("address") != c.Claims().AddressHex {
				return echo.ErrNotFound
			}
			return h(ec)
		}
	}
}
