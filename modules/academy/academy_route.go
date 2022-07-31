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
	}
}
