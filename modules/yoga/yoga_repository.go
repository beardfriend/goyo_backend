package yoga

import (
	"sync"

	"goyo/models/yoga"
	"goyo/server/mariadb"
)

type Repo interface {
	GetList(naverId uint, result *[]YogaDAO) error
}

type repo struct{}

func (repo) GetList(naverId uint, result *[]YogaDAO) error {
	return mariadb.GetInstance().Debug().Model(&yoga.YogaSort{}).Where(&yoga.YogaSort{NaverBasicInfoID: naverId}).Find(&result).Error
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
