package middlewares

import (
	"goyo/libs"
	"goyo/modules/common"

	"github.com/gin-gonic/gin"
)

func ValidateAPIkey(c *gin.Context) {
	APIKey := c.Request.Header.Get("X-API-Key")

	if APIKey == libs.ENV.API.Goyo {
		c.Next()
		return
	}
	common.SendError(c, 401, "API KEY 인증을 다시 확인해주세요")
}
