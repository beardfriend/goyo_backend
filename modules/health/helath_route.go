package health

import (
	"goyo/modules/common"

	"github.com/labstack/echo/v4"
)

func (c HealthController) Routes() []common.Route {
	return []common.Route{
		{
			Method:  echo.POST,
			Path:    "/health",
			Handler: c.Post,
		},
		{
			Method:  echo.GET,
			Path:    "/health",
			Handler: c.Check,
		},
	}
}
