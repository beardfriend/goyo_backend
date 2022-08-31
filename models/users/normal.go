package users

type Normal struct {
	Users    Users `gorm:"foreignkey:Id"`
	Id       uint
	Email    string `gorm:"column:email; size:100; NOT NULL"`
	Password string `gorm:"column:password; size50 NOT NULL"`
}

func (Normal) TableName() string {
	return "users_normal"
}
