package academy

import (
	"goyo/models"
	"goyo/server/mariadb"

	"github.com/labstack/echo/v4"
)

type AcademyController struct{}

func (AcademyController) CrawlNaver(c echo.Context) error {
	var result []models.AdminiStrationDivision
	mariadb.GetInstance().Raw(`SELECT si_gun_gu FROM administration_division WHERE si_do='서울'`).Find(&result)

	return c.JSON(200, result)
}
