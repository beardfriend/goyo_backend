package academy

import (
	"strings"
	"sync"

	"goyo/models/academy"
	"goyo/server/mariadb"

	"gorm.io/gorm/clause"
)

type Repo interface {
	FindNaverBasicInfo(naverId string) int64
	InsertNaverBasicInfo(value *academy.NaverBasicInfo) error
	GetAcademyListByYoga(query *GetListQuery, result *[]NaverBasicInfoDTO) error
	GetAcademyTotalByYoga(query *GetListQuery, total *int64) error
}

type repo struct{}

func (repo) FindNaverBasicInfo(naverId string) int64 {
	var result academy.NaverBasicInfo
	return mariadb.GetInstance().Where("naver_id = ?", naverId).Find(&result).RowsAffected
}

func (repo) InsertNaverBasicInfo(value *academy.NaverBasicInfo) error {
	return mariadb.GetInstance().Create(&value).Error
}

func (repo) GetAcademyTotalByYoga(query *GetListQuery, total *int64) error {
	clauses := make([]clause.Expression, 0)
	if query.SiGunGu != "" {
		clauses = append(clauses, clause.Like{Column: "a.common_address", Value: "%" + query.SiGunGu + "%"})
	}

	if query.YogaSort != "" {
		if strings.Contains(query.YogaSort, ",") {
			ss := strings.Split(query.YogaSort, ",")
			s := make([]interface{}, len(ss))
			for i, v := range ss {
				s[i] = v
			}
			clauses = append(clauses, clause.IN{Column: "b.name", Values: s})
		} else {
			clauses = append(clauses, clause.Eq{Column: "b.name", Value: query.YogaSort})
		}
	}

	return mariadb.GetInstance().
		Debug().
		Clauses(clauses...).
		Table("academy_naver_basic_info a").
		Select("count(a.id) as total").
		Joins("JOIN yoga_sort b ON a.id = b.naver_basic_info_id").
		Group("a.id").
		Count(total).Error
}

func (repo) GetAcademyListByYoga(query *GetListQuery, result *[]NaverBasicInfoDTO) error {
	clauses := make([]clause.Expression, 0)
	offset := 0
	limit := 10

	if query.RowCount != 10 && query.RowCount != 0 {
		limit = query.RowCount
	}

	if query.PageNo > 1 {
		offset = offset + ((query.PageNo - 1) * limit)
	}

	clauses = append(clauses, clause.Limit{Limit: limit, Offset: offset})

	if query.SiGunGu != "" {
		clauses = append(clauses, clause.Like{Column: "a.common_address", Value: "%" + query.SiGunGu + "%"})
	}

	if query.YogaSort != "" {
		if strings.Contains(query.YogaSort, ",") {
			ss := strings.Split(query.YogaSort, ",")
			s := make([]interface{}, len(ss))
			for i, v := range ss {
				s[i] = v
			}
			clauses = append(clauses, clause.IN{Column: "b.name", Values: s})
		} else {
			clauses = append(clauses, clause.Eq{Column: "b.name", Value: query.YogaSort})
		}
	}
	return mariadb.GetInstance().
		Debug().
		Clauses(clauses...).
		Preload("YogaSorts").
		Table("academy_naver_basic_info a").
		Select("a.*").
		Joins("JOIN yoga_sort b ON a.id = b.naver_basic_info_id").
		Group("a.id").
		Find(&result).Error
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
