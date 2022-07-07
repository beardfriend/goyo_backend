package health

import "github.com/labstack/echo/v4"

func Post(c echo.Context) error {
	GetHealthRepo().Insert()
	return c.NoContent(200)
}

func Check(c echo.Context) error {
	result := GetHealthRepo().Get()
	return c.JSON(200, result)
}
