package schedule

import (
	"sync"

	"goyo/models/schedules"
	"goyo/server/mariadb"
)

type Repo interface{}

type repo struct{}

func (repo) Post(model *schedules.TimeTable) error {
	return mariadb.GetInstance().Create(&model).Error
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
