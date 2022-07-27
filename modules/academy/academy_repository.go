package academy

import (
	"sync"

	"goyo/models/academy"
	"goyo/server/mariadb"
)

type AcademyRepo interface {
	FindNaverBasicInfo(naverId string) int64
	InsertNaverBasicInfo(value *academy.NaverBasicInfo) error
}

type academyRepo struct{}

func (academyRepo) FindNaverBasicInfo(naverId string) int64 {
	var result academy.NaverBasicInfo
	return mariadb.GetInstance().Debug().Where("naver_id = ?", naverId).Find(&result).RowsAffected
}

func (academyRepo) InsertNaverBasicInfo(value *academy.NaverBasicInfo) error {
	return mariadb.GetInstance().Debug().Create(&value).Error
}

var (
	academyRepoInstance AcademyRepo
	repoOnce            sync.Once
)

func GetAcademyRepo() AcademyRepo {
	if academyRepoInstance != nil {
		return academyRepoInstance
	}
	repoOnce.Do(func() {
		academyRepoInstance = &academyRepo{}
	})
	return academyRepoInstance
}
