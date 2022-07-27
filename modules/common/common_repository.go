package common

import (
	"sync"

	"goyo/models"
	"goyo/server/mariadb"
)

type CommonRepo interface {
	Getgu(result *[]models.AdminiStrationDivision) error
}

type commonRepo struct{}

func (commonRepo) Getgu(result *[]models.AdminiStrationDivision) error {
	return mariadb.GetInstance().Find(&result).Error
}

var (
	commonRepoInstance CommonRepo
	repoOnce           sync.Once
)

func GetCommonRepo() CommonRepo {
	if commonRepoInstance != nil {
		return commonRepoInstance
	}
	repoOnce.Do(func() {
		commonRepoInstance = &commonRepo{}
	})
	return commonRepoInstance
}
