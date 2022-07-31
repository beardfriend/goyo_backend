package health

import (
	"goyo/models"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (HealthController) Post(c *gin.Context) {
	GetHealthRepo().Insert()
	c.String(200, "")
}

func (HealthController) Check(c *gin.Context) {
	var models models.Health
	GetHealthRepo().Get(&models)

	response := HealthResult(models)
	c.JSON(200, response)
}
