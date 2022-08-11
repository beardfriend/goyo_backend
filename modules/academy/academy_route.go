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
			Path:    "/v1/academy/list",
			Handler: []gin.HandlerFunc{c.GetList},
		},
		{
			Method:  "PUT",
			Path:    "/academy/thumb_url",
			Handler: []gin.HandlerFunc{c.UpdateThumbUrl},
		},
	}
}
