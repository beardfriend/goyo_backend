package auth

import (
	"sync"

	"goyo/interfaces"
	"goyo/models/users"
	"goyo/server/mariadb"
)

type AuthRepo interface {
	Get() (r interfaces.HealthResult)
	Insert()
}

type authRepo struct{}

var (
	singleton AuthRepo
	once      sync.Once
)

func GetAuthRepo() AuthRepo {
	once.Do(func() {
		singleton = &authRepo{}
	})
	return singleton
}

func (authRepo) Insert() {
	mariadb.GetInstance().Create(&users.SocialLogin{Type: "kakao", Token: "asdasd", User: users.User{}})
	// User:mariadb.GetInstance().Create(&users.User{Email: "asd"}).Association()
	// mariadb.GetInstance().Create(&users.SocialLogin{Type: "kakao", Token: "asdasd"})
	// mariadb.GetInstance().Save(&users.SocialLogin{User: users.User{}, Type: "kakao", Token: "asdasd"})
}

func (authRepo) Get() (r interfaces.HealthResult) {
	var result interfaces.HealthResult

	if err := mariadb.GetInstance().First(&result).Error; err != nil {
		panic(err)
	}
	return result
}

//func SetAuthRepo(repo AuthRepo) AuthRepo {
//	original := singleton
//	singleton = repo
//	return original
//}
