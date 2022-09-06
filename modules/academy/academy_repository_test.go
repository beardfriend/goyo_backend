package academy

import (
	"fmt"
	"testing"

	"goyo/libs"
)

func BenchmarkGetAcademyList(b *testing.B) {
	libs.ReadEnv("test")

	for i := 0; i < b.N; i++ {
		query := new(GetListQuery)
		var academyList []NaverPlaceDTO
		query.YogaSort = "하타"
		GetRepo().GetAcademyListByYoga(query, &academyList)
	}
}

func BenchmarkGetTotal(b *testing.B) {
	libs.ReadEnv("test")
	for i := 0; i < b.N; i++ {
		query := new(GetListQuery)
		var total int64
		query.YogaSort = "하타"
		GetRepo().GetAcademyTotalByYoga(query, &total)
	}
}

func TestGetAcademyList(t *testing.T) {
	libs.ReadEnv("test")
	query := new(GetListQuery)
	var academyList []NaverPlaceDTO
	query.YogaSort = "하타"
	GetRepo().GetAcademyListByYoga(query, &academyList)
	fmt.Println(academyList[0].BookingUrl)
}
