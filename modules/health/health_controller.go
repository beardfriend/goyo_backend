package health

import (
	"goyo/models"

	"github.com/labstack/echo/v4"
)

type HealthController struct{}

func (HealthController) Post(c echo.Context) error {
	GetHealthRepo().Insert()
	return c.NoContent(200)
}

func (HealthController) Check(c echo.Context) error {
	var models models.Health
	GetHealthRepo().Get(&models)

	response := HealthResult(models)
	return c.JSON(200, response)
}
