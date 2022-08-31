package users

type Normal struct {
	Users    Users `gorm:"foreignkey:Id"`
	Id       uint
	Password string `gorm:"size50"`
}

func (Normal) TableName() string {
	return "users_normal"
}
