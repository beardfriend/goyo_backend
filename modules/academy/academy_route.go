package academy

import (
	"goyo/modules/common"

	"github.com/gin-gonic/gin"
)

func (c AcademyController) Routes() []common.Route {
	return []common.Route{
		{
			Method:  "GET",
			Path:    "/academy",
			Handler: []gin.HandlerFunc{c.CrawlNaver},
		},
		{
			Method:  "GET",
			Path:    "/academy/list",
			Handler: []gin.HandlerFunc{c.GetList},
		},
	}
}
