package yoga

import (
	"sync"

	"goyo/models/yoga"
	"goyo/server/mariadb"
)

type Repo interface {
	GetYogaSortByName(name string, result *[]YogaSorts) error
	GetYogaSortByCosonants(firstWord string, lastWord string, result *[]YogaSorts) error
}

type repo struct{}

func (repo) GetYogaSortByName(name string, result *[]YogaSorts) error {
	return mariadb.GetInstance().Debug().Select("distinct(name)").Model(&yoga.YogaSort{}).Group("name").Where("name LIKE ?", name+"%").Limit(6).Find(&result).Error
}

func (repo) GetYogaSortDistinct(result *[]YogaSorts) error {
	return mariadb.GetInstance().Debug().Select("distinct(name)").Model(&yoga.YogaSort{}).Find(&result).Error
}

func (repo) GetYogaSortByCosonants(firstWord string, lastWord string, result *[]YogaSorts) error {
	return mariadb.GetInstance().Debug().Select("distinct(name)").Model(&yoga.YogaSort{}).Group("name").Where("name >= ? AND name <= ?", firstWord, lastWord).Limit(6).Find(&result).Error
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
