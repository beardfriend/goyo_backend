package yoga

import (
	"goyo/modules/common"

	"github.com/gin-gonic/gin"
)

type YogaController struct{}

func (YogaController) GET(c *gin.Context) {
	var result []YogaSorts
	if err := GetRepo().GetYogaSort(&result); err != nil {
		panic(err)
	}

	var response []GetResponse

	for i, v := range result {
		response = append(response, GetResponse{Id: uint(i + 1), Name: v.Name})
	}

	common.SendResult(c, 200, "성공적으로 조회했습니다.", response)
}
