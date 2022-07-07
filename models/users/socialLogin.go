package users

type SocialLogin struct {
	User  User `gorm:"foreignkey:Id"`
	Id    uint
	Type  string `gorm:"size:10"`
	Token string `gorm:"size:100"`
}
