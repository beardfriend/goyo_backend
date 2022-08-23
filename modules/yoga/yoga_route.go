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
		{
			Method:  "GET",
			Path:    "/v2/yoga/sorts",
			Handler: []gin.HandlerFunc{c.GetSortsV2},
		},
		{
			Method:  "POST",
			Path:    "/yoga/sorts",
			Handler: []gin.HandlerFunc{c.CronYogaSorts},
		},
		{
			Method:  "PUT",
			Path:    "/yoga/sorts/score",
			Handler: []gin.HandlerFunc{c.UpdateScore},
		},
		{
			Method:  "POST",
			Path:    "/yoga/sorts/score",
			Handler: []gin.HandlerFunc{c.Ranking},
		},
		{
			Method:  "GET",
			Path:    "/yoga/sorts/score",
			Handler: []gin.HandlerFunc{c.GetRanking},
		},
	}
}
