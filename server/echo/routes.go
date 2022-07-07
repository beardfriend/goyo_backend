package echo

import (
	"goyo/modules/auth"
	"goyo/modules/health"

	"github.com/labstack/echo/v4"
)

func Routes(c *echo.Echo) {
	healthGroup := c.Group("health")
	{
		healthGroup.GET("", health.Check)
		healthGroup.GET("/update", health.Post)
	}

	authG := c.Group("auth")
	{
		authG.GET("", auth.Post)
	}
}
