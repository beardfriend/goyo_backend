package common

import (
	"github.com/gin-gonic/gin"
)

func (c CommonController) Routes() []Route {
	return []Route{
		{
			Method:  "GET",
			Path:    "/administrations",
			Handler: []gin.HandlerFunc{c.GetAdministrations},
		},
	}
}
