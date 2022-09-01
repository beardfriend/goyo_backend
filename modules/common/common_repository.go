package common

import (
	"sync"

	"goyo/models"
	"goyo/server/mariadb"

	"gorm.io/gorm/clause"
)

type CommonRepo interface {
	GetAdminiStrationDivision(result *[]models.AdminiStrationDivision) error
	CommonListClauses(pageNo, rowCount int) clause.Expression
}

type commonRepo struct{}

// ------------------- AdminiStration -------------------

func (commonRepo) GetAdminiStrationDivision(result *[]models.AdminiStrationDivision) error {
	return mariadb.GetInstance().Find(&result).Error
}

// ------------------- Common List Clauses -------------------

func (commonRepo) CommonListClauses(pageNo, rowCount int) clause.Expression {
	offset := 0
	limit := 10

	if rowCount != 10 && rowCount != 0 {
		limit = rowCount
	}

	if pageNo > 1 {
		offset = offset + ((pageNo - 1) * limit)
	}

	return clause.Limit{Limit: limit, Offset: offset}
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
