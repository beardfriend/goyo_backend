package models

type AdminiStrationDivision struct {
	SiGunGu string `json:"si_gun_gu"`
}

func (AdminiStrationDivision) TableName() string {
	return `administration_division`
}
