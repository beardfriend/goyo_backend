package yoga

import (
	"sync"

	"goyo/models/yoga"
	"goyo/server/mariadb"
)

type Repo interface {
	GetYogaSort(result *[]YogaSorts) error
}

type repo struct{}

func (repo) GetYogaSort(result *[]YogaSorts) error {
	return mariadb.GetInstance().Select("distinct(name)").Debug().Model(&yoga.YogaSort{}).Group("name").Find(&result).Error
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
