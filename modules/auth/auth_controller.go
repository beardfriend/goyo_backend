package auth

import "github.com/labstack/echo/v4"

func Post(c echo.Context) error {
	GetAuthRepo().Insert()
	return c.NoContent(200)
}
