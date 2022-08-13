package academy

import (
	"goyo/modules/common"
	"goyo/modules/middlewares"

	"github.com/gin-gonic/gin"
)

func (c AcademyController) Routes() []common.Route {
	return []common.Route{
		{
			Method:  "GET",
			Path:    "/academy",
			Handler: []gin.HandlerFunc{middlewares.ValidateAPIkey, c.CrawlNaver},
		},
		{
			Method:  "GET",
			Path:    "/v1/academy/list",
			Handler: []gin.HandlerFunc{c.GetList},
		},
		{
			Method:  "PUT",
			Path:    "/academy/thumb_url",
			Handler: []gin.HandlerFunc{middlewares.ValidateAPIkey, c.UpdateThumbUrl},
		},
	}
}
