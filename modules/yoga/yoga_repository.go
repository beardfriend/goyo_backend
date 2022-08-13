package yoga

import (
	"sync"

	"goyo/models/yoga"
	"goyo/server/mariadb"
)

type Repo interface {
	GetSortsByName(name string, result *[]SortsDTO) error
	GetSortsByCosonants(firstWord string, lastWord string, result *[]SortsDTO) error
}

type repo struct{}

func (repo) GetSortsByName(name string, result *[]SortsDTO) error {
	return mariadb.GetInstance().Debug().Select("distinct(name)").Model(&yoga.YogaSorts{}).Group("name").Where("name LIKE ?", name+"%").Limit(6).Find(&result).Error
}

func (repo) GetYogaSortDistinct(result *[]SortsDTO) error {
	return mariadb.GetInstance().Debug().Select("distinct(name)").Model(&yoga.YogaSorts{}).Find(&result).Error
}

func (repo) GetSortsByCosonants(firstWord string, lastWord string, result *[]SortsDTO) error {
	return mariadb.GetInstance().Debug().Select("distinct(name)").Model(&yoga.YogaSorts{}).Group("name").Where("name >= ? AND name <= ?", firstWord, lastWord).Limit(6).Find(&result).Error
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
