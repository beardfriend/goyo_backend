package yoga

import (
	"sync"

	"goyo/models/yoga"
	"goyo/server/mariadb"
)

type Repo interface {
	GetSortsByName(name string, result *[]SortsDTO) error
	GetSortsByCosonants(firstWord string, lastWord string, result *[]SortsDTO) error
	CreateSorts(value *[]CreateSortsDTO) error
}

type repo struct{}

// ------------------- Get -------------------

func (repo) GetSortsByName(name string, result *[]SortsDTO) error {
	return mariadb.GetInstance().
		Select("distinct(name)").
		Model(&yoga.YogaSorts{}).
		Where("name LIKE ?", name+"%").
		Group("name").
		Limit(8).
		Order("RAND ()").
		Find(&result).Error
}

func (repo) GetYogaSortDistinct(result *[]SortsDTO) error {
	return mariadb.GetInstance().
		Select("distinct(name)").
		Model(&yoga.YogaSorts{}).
		Find(&result).Error
}

func (repo) GetSortsByCosonants(firstWord string, lastWord string, result *[]SortsDTO) error {
	return mariadb.GetInstance().
		Select("distinct(name)").
		Model(&yoga.YogaSorts{}).
		Where("name >= ? AND name <= ?", firstWord, lastWord).
		Group("name").
		Limit(8).
		Order("RAND ()").
		Find(&result).Error
}

// ------------------- Create -------------------

func (repo) CreateSorts(value *[]CreateSortsDTO) error {
	return mariadb.GetInstance().
		Create(&value).Error
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
