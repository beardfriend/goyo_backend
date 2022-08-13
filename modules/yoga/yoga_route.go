package yoga

import (
	"goyo/modules/common"

	"github.com/gin-gonic/gin"
)

func (c YogaController) Routes() []common.Route {
	return []common.Route{
		{
			Method:  "GET",
			Path:    "/v1/yoga/category",
			Handler: []gin.HandlerFunc{c.GetSorts},
		},
		{
			Method:  "POST",
			Path:    "/v1/yoga/category",
			Handler: []gin.HandlerFunc{c.InsertYoga},
		},
	}
}
