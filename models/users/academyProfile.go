package users

import "goyo/models"

type AcademyProfile struct {
	models.Primary
	Name string
	models.DefaultTime
}

func (AcademyProfile) TableName() string {
	return `academies_profile`
}
