package admin

import (
	"sync"

	"goyo/models/naver"
	"goyo/models/yoga"
	"goyo/modules/academy"
	"goyo/server/mariadb"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repo interface {
	GetListThatHasntTag(query *academy.AcademyListRequest, result *[]NaverPlaceDTO) error
	GetListThatHasnTagTotal(query *academy.AcademyListRequest, result *int64) error
	GetDetail(naverId *uint64, result *GetDetailNaverPlaceDTO) error
	DeleteSorts(idList *[]uint64) error
}

type repo struct{}

var (
	repoInstance Repo
	repoOnce     sync.Once
)

func (repo) GetListThatHasntTag(query *academy.AcademyListRequest, result *[]NaverPlaceDTO) error {
	clauses := make([]clause.Expression, 0)

	offset := 0
	limit := 10

	if query.RowCount != 10 && query.RowCount != 0 {
		limit = query.RowCount
	}

	if query.PageNo > 1 {
		offset = offset + ((query.PageNo - 1) * limit)
	}

	if query.SiGunGu != "" {
		clauses = append(clauses, clause.Like{Column: "a.common_address", Value: "%" + query.SiGunGu + "%"})
	}

	if !query.ContainMeditation {
		clauses = append(clauses, clause.Eq{Column: "a.category", Value: "요가원"})
	}

	if query.Status == "Regist" {
		return mariadb.GetInstance().
			Select("a.*").
			Debug().
			Table("naver_place a").
			Joins("INNER JOIN yoga_sorts b ON a.id = b.naver_place_id").
			Clauses(clauses...).
			Group("a.id").
			Limit(limit).
			Offset(offset).
			Find(&result).Error
	}

	if query.Status == "NonRegist" {
		return mariadb.GetInstance().
			Select("a.*").
			Debug().
			Table("naver_place a").
			Joins("LEFT JOIN yoga_sorts b ON a.id = b.naver_place_id").
			Where("b.naver_place_id IS NULL").
			Clauses(clauses...).
			Group("a.id").
			Limit(limit).
			Offset(offset).
			Find(&result).Error
	}

	return mariadb.GetInstance().
		Select("a.*, IF( (SELECT count(id) FROM yoga_sorts WHERE naver_place_id = a.id ) > 0 , 1, 0) as is_regist").
		Debug().
		Table("naver_place a").
		Clauses(clauses...).
		Group("a.id").
		Limit(limit).
		Offset(offset).
		Find(&result).Error
}

func (repo) GetListThatHasnTagTotal(query *academy.AcademyListRequest, result *int64) error {
	clauses := make([]clause.Expression, 0)

	if query.SiGunGu != "" {
		clauses = append(clauses, clause.Like{Column: "a.common_address", Value: "%" + query.SiGunGu + "%"})
	}

	if !query.ContainMeditation {
		clauses = append(clauses, clause.Eq{Column: "a.category", Value: "요가원"})
	}

	if query.Status == "Regist" {
		return mariadb.GetInstance().
			Debug().
			Table("naver_place a").
			Joins("INNER JOIN yoga_sorts b ON a.id = b.naver_place_id").
			Group("a.id").
			Clauses(clauses...).
			Count(result).Error
	}

	if query.Status == "NonRegist" {
		return mariadb.GetInstance().
			Debug().
			Table("naver_place a").
			Joins("LEFT JOIN yoga_sorts b ON a.id = b.naver_place_id").
			Where("b.naver_place_id IS NULL").
			Group("a.id").
			Clauses(clauses...).
			Count(result).Error
	}
	return mariadb.GetInstance().
		Debug().
		Table("naver_place a").
		Group("a.id").
		Clauses(clauses...).
		Count(result).Error
}

func (repo) GetDetail(naverId *uint64, result *GetDetailNaverPlaceDTO) error {
	return mariadb.GetInstance().
		Preload("YogaSorts", func(db *gorm.DB) *gorm.DB {
			return db.Order("yoga_sorts.name")
		}).
		Model(&naver.NaverPlace{}).
		Where("id = ?", naverId).
		Find(&result).Error
}

func (repo) DeleteSorts(idList *[]uint64) error {
	return mariadb.GetInstance().Debug().
		Delete(&yoga.YogaSorts{}, idList).Error
}

func GetRepo() Repo {
	if repoInstance != nil {
		return repoInstance
	}
	repoOnce.Do(func() {
		repoInstance = &repo{}
	})
	return repoInstance
}
