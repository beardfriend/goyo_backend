package users

type Kakao struct {
	Social   Social `gorm:"foreignkey:SocialId"`
	SocialId uint   `gorm:"primaryKey"`
	Email    string `gorm:"size40"`
}

func (Kakao) TableName() string {
	return "users_social_kakao"
}
