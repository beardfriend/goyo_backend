package health

import (
	"sync"

	"goyo/models"
	"goyo/server/mariadb"
)

type HealthRepo interface {
	Get(r *models.Health)
	Insert(value string)
}

type healthRepo struct{}

var (
	singleton HealthRepo
	once      sync.Once
)

func GetHealthRepo() HealthRepo {
	if singleton != nil {
		return singleton
	}
	once.Do(func() {
		singleton = &healthRepo{}
	})
	return singleton
}

func (healthRepo) Insert(value string) {
	if err := mariadb.GetInstance().Save(&models.Health{Status: value}).Error; err != nil {
		panic(err)
	}
}

func (healthRepo) Get(r *models.Health) {
	if err := mariadb.GetInstance().First(&r).Error; err != nil {
		panic(err)
	}
}

//func SetHealthRepo(repo HealthRepo) HealthRepo {
//	original := singleton
//	singleton = repo
//	return original
//}
