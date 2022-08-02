package academy

import (
	"sync"

	"goyo/models/academy"
	"goyo/server/mariadb"
)

type Repo interface {
	FindNaverBasicInfo(naverId string) int64
	InsertNaverBasicInfo(value *academy.NaverBasicInfo) error
	GetAcademyListByYoga(sort string, result *[]NaverBasicInfoDAO) error
}

type repo struct{}

func (repo) FindNaverBasicInfo(naverId string) int64 {
	var result academy.NaverBasicInfo
	return mariadb.GetInstance().Where("naver_id = ?", naverId).Find(&result).RowsAffected
}

func (repo) InsertNaverBasicInfo(value *academy.NaverBasicInfo) error {
	return mariadb.GetInstance().Create(&value).Error
}

func (repo) GetAcademyListByYoga(sort string, result *[]NaverBasicInfoDAO) error {
	return mariadb.GetInstance().
		Debug().
		Table("academy_naver_basic_info a").
		Select("a.*").
		Joins("JOIN yoga_sort b ON a.id = b.naver_basic_info_id").
		Where("b.name = ?", sort).
		Find(&result).Error
}

var (
	repoInstance Repo
	repoOnce     sync.Once
)

func GetRepo() Repo {
	if repoInstance != nil {
		return repoInstance
	}
	repoOnce.Do(func() {
		repoInstance = &repo{}
	})
	return repoInstance
}
