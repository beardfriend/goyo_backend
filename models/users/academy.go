package users

import "goyo/models"

type Acadmey struct {
	models.Primary
	AcademyProfileID int            `gorm:"column:acadmey_profile_id"`
	AcademyProfile   AcademyProfile `gorm:"foreignKey:acadmey_profile_id"`
	Name             string
}

func (Acadmey) TableName() string {
	return `academies`
}
