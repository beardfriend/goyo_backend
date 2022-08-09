package health

import (
	"goyo/models"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (HealthController) Post(c *gin.Context) {
	go Rot()
	c.String(200, "")
}

func (HealthController) Check(c *gin.Context) {
	var models models.Health
	GetHealthRepo().Get(&models)

	response := HealthResult(models)
	c.JSON(200, response)
}

func Rot() {
	arr := []string{"d", "s", "f", "j", "k", "l"}
	for i := 0; i < len(arr); i++ {
		GetHealthRepo().Insert(arr[i])
	}
}
