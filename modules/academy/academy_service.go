package academy

import (
	"sync"

	"goyo/modules/common"
	"goyo/modules/yoga"
)

type Serivce interface {
	NewGetListResponse(result []NaverBasicInfoDAO, yogaSort [][]yoga.YogaDAO, total int64, query *GetListQuery) GetListResponse
	NewYogaList(value []NaverBasicInfoDAO) [][]yoga.YogaDAO
}

type service struct{}

func (service) NewYogaList(value []NaverBasicInfoDAO) [][]yoga.YogaDAO {
	var result [][]yoga.YogaDAO
	for i := 0; i < len(value); i++ {
		var yogaSort []yoga.YogaDAO
		if err := yoga.GetRepo().GetList(value[i].Id, &yogaSort); err != nil {
			panic(err)
		}
		result = append(result, yogaSort)
	}
	return result
}

func (service) NewGetListResponse(result []NaverBasicInfoDAO, yogaSort [][]yoga.YogaDAO, total int64, query *GetListQuery) GetListResponse {
	var response GetListResponse
	list := make([]GetListDetail, 0)
	for i, v := range result {
		sort := make([]yoga.Response, 0)
		for _, j := range yogaSort[i] {
			sort = append(sort, yoga.Response{ID: j.Id, Name: j.Name})
		}
		list = append(list, GetListDetail{
			Id:            v.Id,
			Name:          v.Name,
			Category:      v.Category,
			RoadAddress:   v.RoadAddress,
			CommonAddress: v.Category,
			BookingUrl:    v.BookingUrl,
			PhoneNum:      v.PhoneNum,
			BusinessHours: v.BusinessHours,
			ImageUrl:      v.ImageUrl,
			X:             v.X,
			Y:             v.Y,
			YogaSort:      sort,
		})
	}

	pageInfo := common.PaginationInfo{
		PageSize:  len(list),
		PageCount: total/int64(query.RowCount) + 1,
		Page:      int64(query.PageNo),
		RowCount:  total,
	}

	response.List = list
	response.Pagination = pageInfo

	return response
}

var (
	serviceInstance Serivce
	serviceOnce     sync.Once
)

func GetService() Serivce {
	if serviceInstance != nil {
		return serviceInstance
	}
	serviceOnce.Do(func() {
		serviceInstance = &service{}
	})
	return serviceInstance
}
