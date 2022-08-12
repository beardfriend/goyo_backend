package users

type Social struct {
	Users  Users `gorm:"foreignkey:Id"`
	Id     uint
	Type   int8   `gorm:"comment:1.카카오 2.구글"`
	Secret string `gorm:"size:100"`
}

func (Social) TableName() string {
	return "users_social"
}
