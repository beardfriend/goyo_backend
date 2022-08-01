package academy

import (
	"sync"

	"goyo/models/academy"
	"goyo/server/mariadb"
)

type Repo interface {
	FindNaverBasicInfo(naverId string) int64
	InsertNaverBasicInfo(value *academy.NaverBasicInfo) error
}

type repo struct{}

func (repo) FindNaverBasicInfo(naverId string) int64 {
	var result academy.NaverBasicInfo
	return mariadb.GetInstance().Where("naver_id = ?", naverId).Find(&result).RowsAffected
}

func (repo) InsertNaverBasicInfo(value *academy.NaverBasicInfo) error {
	return mariadb.GetInstance().Create(&value).Error
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
