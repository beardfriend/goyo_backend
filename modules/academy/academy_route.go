package academy

import (
	"goyo/modules/common"

	"github.com/labstack/echo/v4"
)

func (c AcademyController) Routes() []common.Route {
	return []common.Route{
		{
			Method:  echo.GET,
			Path:    "/academy",
			Handler: c.CrawlNaver,
		},
	}
}
