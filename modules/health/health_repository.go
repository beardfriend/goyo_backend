package health

import (
	"sync"

	"goyo/interfaces"
	"goyo/models"
	"goyo/server/mariadb"
)

type HealthRepo interface {
	Get() (r interfaces.HealthResult)
	Insert()
}

type healthRepo struct{}

var (
	singleton HealthRepo
	once      sync.Once
)

func GetHealthRepo() HealthRepo {
	once.Do(func() {
		singleton = &healthRepo{}
	})
	return singleton
}

func (healthRepo) Insert() {
	if err := mariadb.GetInstance().Save(&models.Health{Status: "ok"}).Error; err != nil {
		panic(err)
	}
}

func (healthRepo) Get() (r interfaces.HealthResult) {
	var result interfaces.HealthResult

	if err := mariadb.GetInstance().First(&result).Error; err != nil {
		panic(err)
	}
	return result
}

//func SetHealthRepo(repo HealthRepo) HealthRepo {
//	original := singleton
//	singleton = repo
//	return original
//}
