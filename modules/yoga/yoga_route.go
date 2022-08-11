package yoga

import (
	"goyo/modules/common"

	"github.com/gin-gonic/gin"
)

func (c YogaController) Routes() []common.Route {
	return []common.Route{
		{
			Method:  "GET",
			Path:    "/yoga/category",
			Handler: []gin.HandlerFunc{c.GET},
		},
		{
			Method:  "POST",
			Path:    "/yoga/category",
			Handler: []gin.HandlerFunc{c.InsertYoga},
		},
	}
}
