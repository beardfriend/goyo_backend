package yoga

import (
	"sync"

	"goyo/models/yoga"
	"goyo/server/mariadb"
)

type Repo interface {
	GetYogaSort(name string, result *[]YogaSorts) error
	GetYogaSortByCosonants(firstWord string, lastWord string, result *[]YogaSorts) error
}

type repo struct{}

func (repo) GetYogaSort(name string, result *[]YogaSorts) error {
	return mariadb.GetInstance().Select("distinct(name)").Model(&yoga.YogaSort{}).Group("name").Where("name LIKE ?", name+"%").Find(&result).Error
}

func (repo) GetYogaSortByCosonants(firstWord string, lastWord string, result *[]YogaSorts) error {
	return mariadb.GetInstance().Debug().Select("distinct(name)").Model(&yoga.YogaSort{}).Group("name").Where("name >= ? AND name <= ?", firstWord, lastWord).Find(&result).Error
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
