package yoga

import (
	"goyo/modules/common"

	"github.com/gin-gonic/gin"
)

func (c YogaController) Routes() []common.Route {
	return []common.Route{
		{
			Method:  "GET",
			Path:    "/v1/yoga/sorts",
			Handler: []gin.HandlerFunc{c.GetSorts},
		},
	}
}
