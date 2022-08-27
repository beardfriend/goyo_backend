package yoga

import (
	"goyo/modules/common"
	"goyo/modules/middlewares"

	"github.com/gin-gonic/gin"
)

func (c YogaController) Routes() []common.Route {
	return []common.Route{
		{
			Method:  "GET",
			Path:    "/v1/yoga/sorts",
			Handler: []gin.HandlerFunc{c.GetSorts},
		},
		{
			Method:  "GET",
			Path:    "/v2/yoga/sorts",
			Handler: []gin.HandlerFunc{c.GetSortsV2},
		},
		{
			Method:  "POST",
			Path:    "/yoga/sorts",
			Handler: []gin.HandlerFunc{middlewares.ValidateAPIkey, c.CronYogaSorts},
		},
		{
			Method:  "PUT",
			Path:    "/yoga/sorts/score",
			Handler: []gin.HandlerFunc{middlewares.ValidateAPIkey, c.UpdateScore},
		},
		{
			Method:  "POST",
			Path:    "/yoga/sorts/score",
			Handler: []gin.HandlerFunc{middlewares.ValidateAPIkey, c.Ranking},
		},
		{
			Method:  "GET",
			Path:    "/yoga/sorts/ranking",
			Handler: []gin.HandlerFunc{c.GetRanking},
		},
	}
}
