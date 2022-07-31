package common

import "github.com/gin-gonic/gin"

func SendResult(c *gin.Context, code int, message string, result interface{}) {
	c.JSON(code, gin.H{"message": message, "result": result})
}

func SendError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{"message": message})
}

func SendOk(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"message": message})
}
