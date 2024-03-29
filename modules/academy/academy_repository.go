package academy

import (
	"strings"
	"sync"

	"goyo/models/naver"
	"goyo/modules/common"

	"goyo/server/mariadb"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repo interface {
	FindNaverPlace(naverId string, result *naver.NaverPlace) (int64, error)
	CreateNaverPlace(newData *naver.NaverPlace) error
	GetAcademyListByYoga(query *GetListQuery, result *[]NaverPlaceDTO) error
	GetAcademyTotalByYoga(query *GetListQuery, total *int64) error
	GetNaverId(result *[]CrawlMobileNaverPlaceDto) error
	UpdateNaverBasicInfo(id uint, thumbUrl string) error
}

type repo struct{}

// ------------------- Find -------------------

func (repo) FindNaverPlace(naverId string, result *naver.NaverPlace) (int64, error) {
	query := mariadb.GetInstance().
		Model(&naver.NaverPlace{}).
		Where("naver_id = ?", naverId).Find(&result)

	return query.RowsAffected, query.Error
}

// ------------------- Get -------------------

func (repo) GetNaverId(result *[]CrawlMobileNaverPlaceDto) error {
	return mariadb.GetInstance().Model(&naver.NaverPlace{}).Where("thumb_url IS NULL").Find(&result).Error
}

func (repo) GetAcademyTotalByYoga(query *GetListQuery, total *int64) error {
	clauses := AcademyListClause(query)

	return mariadb.GetInstance().
		Clauses(clauses...).
		Table("naver_place a").
		Select("count(a.id) as total").
		Joins("JOIN yoga_sorts b ON a.id = b.naver_place_id").
		Group("a.id").
		Count(total).Error
}

func (repo) GetAcademyListByYoga(query *GetListQuery, result *[]NaverPlaceDTO) error {
	listClauses := common.GetCommonRepo().CommonListClauses(query.PageNo, query.RowCount)
	clauses := AcademyListClause(query)
	clauses = append(clauses, listClauses)

	return mariadb.GetInstance().
		Clauses(clauses...).
		Preload("YogaSorts", func(db *gorm.DB) *gorm.DB {
			c := make([]clause.Expression, 0)
			c = append(c, clause.OrderBy{Expression: clause.Expr{
				SQL:  "(CASE WHEN name LIKE ? THEN 1 WHEN name LIKE ? THEN 2 ELSE 3 END), name ASC",
				Vars: []interface{}{query.YogaSort, "%" + query.YogaSort + "%"},
			}})
			return db.Clauses(c...)
		}).
		Table("naver_place a").
		Select("a.*").
		Joins("JOIN yoga_sorts b ON a.id = b.naver_place_id").
		Group("a.id").
		Find(&result).Error
}

// ------------------- Create -------------------

func (repo) CreateNaverPlace(newData *naver.NaverPlace) error {
	return mariadb.GetInstance().Create(&newData).Error
}

// ------------------- Update -------------------

func (repo) UpdateNaverBasicInfo(id uint, thumbUrl string) error {
	return mariadb.GetInstance().Model(&naver.NaverPlace{}).Where("id = ?", id).Update("thumb_url", thumbUrl).Error
}

// ------------------- Clasuses -------------------

func AcademyListClause(query *GetListQuery) []clause.Expression {
	clauses := make([]clause.Expression, 0)
	if query.SiGunGu != "" {
		clauses = append(clauses, clause.Like{Column: "a.common_address", Value: "%" + query.SiGunGu + "%"})
	}

	if query.YogaSort != "" {
		if strings.Contains(query.YogaSort, ",") {
			ss := strings.Split(query.YogaSort, ",")
			arr := make([]clause.Expression, 0)
			for _, v := range ss {
				arr = append(arr, clause.Like{Column: "b.name", Value: "%" + v + "%"})
			}

			clauses = append(clauses, clause.Or(arr...))
		} else {
			clauses = append(clauses, clause.Like{Column: "b.name", Value: "%" + query.YogaSort + "%"})
		}
	}

	return clauses
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
