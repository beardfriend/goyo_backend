package users

type Google struct {
	Social   Social `gorm:"foreignkey:SocialId"`
	SocialId uint   `gorm:"primaryKey"`
	Email    string `gorm:"size40"`
}

func (Google) TableName() string {
	return "users_social_google"
}
