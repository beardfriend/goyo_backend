package health

import (
	"goyo/modules/common"

	"github.com/gin-gonic/gin"
)

func (c HealthController) Routes() []common.Route {
	return []common.Route{
		{
			Method:  "GET",
			Path:    "/health",
			Handler: []gin.HandlerFunc{c.Check},
		},
	}
}
