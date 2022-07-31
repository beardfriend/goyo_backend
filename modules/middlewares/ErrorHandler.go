package middlewares

import (
	"goyo/modules/common"

	"github.com/gin-gonic/gin"
)

func ErrorHandleRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		switch recovered.(type) {
		default:
			common.SendError(c, 500, "일시적인 에러가 발생했습니다.")
		}
	})
}
