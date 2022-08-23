package yoga

import (
	"sync"

	"goyo/models/yoga"
	"goyo/server/mariadb"
)

type Repo interface {
	GetSortsByName(name string, result *[]SortsDTO) error
	GetYogaSortDistinct(result *[]SortsDTO) error
	GetSortsByCosonants(firstWord string, lastWord string, result *[]SortsDTO) error
	CreateSorts(value *[]CreateSortsDTO) error
	CreateCounts(value *[]yoga.YogaScore) error
	GetScores(names []string, result *[]yoga.YogaScore) error
	GetRanking(result *[]yoga.YogaScore) error
	UpdateCounts(id uint, score uint) error
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

func (repo) GetScores(names []string, result *[]yoga.YogaScore) error {
	return mariadb.GetInstance().Where("name IN ?", names).Find(&result).Error
}

func (repo) GetRanking(result *[]yoga.YogaScore) error {
	return mariadb.GetInstance().Order("score DESC").Limit(10).Find(&result).Error
}

// ------------------- Create -------------------

func (repo) CreateSorts(value *[]CreateSortsDTO) error {
	return mariadb.GetInstance().Create(&value).Error
}

func (repo) CreateCounts(value *[]yoga.YogaScore) error {
	return mariadb.GetInstance().Create(&value).Error
}

// ------------------- Update -------------------

func (repo) UpdateCounts(id uint, score uint) error {
	return mariadb.GetInstance().Model(&yoga.YogaScore{}).Where("id = ?", id).Update("score", score).Error
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
