package schedule

import (
	"sync"

	"goyo/models/schedule"
	"goyo/server/mariadb"
)

type Repo interface{}

type repo struct{}

func (repo) Post(model *schedule.Schedule) error {
	return mariadb.GetInstance().Create(&model).Error
}

// ------------------- SingleTon -------------------

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
