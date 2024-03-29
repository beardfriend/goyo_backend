package admin

import (
	"goyo/modules/common"
	"goyo/modules/middlewares"

	"github.com/gin-gonic/gin"
)

func (c AdminController) Routes() []common.Route {
	return []common.Route{
		{
			Method:  "GET",
			Path:    "/admin/academies",
			Handler: []gin.HandlerFunc{middlewares.ValidateAPIkey, c.GetAcademies},
		},
		{
			Method:  "POST",
			Path:    "/admin/yoga/sorts",
			Handler: []gin.HandlerFunc{middlewares.ValidateAPIkey, c.InsertYogaSorts},
		},
		{
			Method:  "GET",
			Path:    "/admin/academy/:naver_id",
			Handler: []gin.HandlerFunc{c.GetAcademyDetail},
		},
		{
			Method:  "DELETE",
			Path:    "/admin/academy/:naver_id",
			Handler: []gin.HandlerFunc{middlewares.ValidateAPIkey, c.DeleteYogaSorts},
		},
	}
}
