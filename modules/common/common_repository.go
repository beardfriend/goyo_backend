package common

import (
	"sync"

	"goyo/models"
	"goyo/server/mariadb"
)

type CommonRepo interface {
	GetAdminiStrationDivision(result *[]models.AdminiStrationDivision) error
}

type commonRepo struct{}

// ------------------- AdminiStration -------------------

func (commonRepo) GetAdminiStrationDivision(result *[]models.AdminiStrationDivision) error {
	return mariadb.GetInstance().Find(&result).Error
}

// ------------------- SINGLETON -------------------
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
