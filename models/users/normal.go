package users

type Normal struct {
	User     User `gorm:"foreignkey:Id"`
	Id       uint
	Email    string `gorm:"size40"`
	Password string `gorm:"size50"`
}

func (Normal) TableName() string {
	return "users_normal"
}
