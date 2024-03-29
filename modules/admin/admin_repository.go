package admin

import (
	"sync"

	"goyo/models/naver"
	"goyo/models/yoga"
	"goyo/modules/academy"
	"goyo/modules/common"
	"goyo/server/mariadb"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repo interface {
	GetAcademiesByRegistered(query *academy.AcademyListRequest, result *[]NaverPlaceDTO) error
	GetAcademieTotalByRegistered(query *academy.AcademyListRequest, result *int64) error
	GetDetail(naverId *uint64, result *GetDetailNaverPlaceDTO) error
	DeleteSorts(idList *[]uint64) error
}

type repo struct{}

var (
	repoInstance Repo
	repoOnce     sync.Once
)

func (repo) GetAcademiesByRegistered(query *academy.AcademyListRequest, result *[]NaverPlaceDTO) error {
	clauses := make([]clause.Expression, 0)
	listClauses := common.GetCommonRepo().CommonListClauses(query.PageNo, query.RowCount)

	clauses = append(clauses, listClauses)
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
			Find(&result).Error
	}

	return mariadb.GetInstance().
		Select("a.*, IF( (SELECT count(id) FROM yoga_sorts WHERE naver_place_id = a.id ) > 0 , 1, 0) as is_regist").
		Table("naver_place a").
		Clauses(clauses...).
		Group("a.id").
		Find(&result).Error
}

func (repo) GetAcademieTotalByRegistered(query *academy.AcademyListRequest, result *int64) error {
	clauses := make([]clause.Expression, 0)

	if query.SiGunGu != "" {
		clauses = append(clauses, clause.Like{Column: "a.common_address", Value: "%" + query.SiGunGu + "%"})
	}

	if !query.ContainMeditation {
		clauses = append(clauses, clause.Eq{Column: "a.category", Value: "요가원"})
	}

	if query.Status == "Regist" {
		return mariadb.GetInstance().
			Table("naver_place a").
			Joins("INNER JOIN yoga_sorts b ON a.id = b.naver_place_id").
			Group("a.id").
			Clauses(clauses...).
			Count(result).Error
	}

	if query.Status == "NonRegist" {
		return mariadb.GetInstance().
			Table("naver_place a").
			Joins("LEFT JOIN yoga_sorts b ON a.id = b.naver_place_id").
			Where("b.naver_place_id IS NULL").
			Group("a.id").
			Clauses(clauses...).
			Count(result).Error
	}
	return mariadb.GetInstance().
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
	return mariadb.GetInstance().
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
