package academy

import (
	"sync"

	"goyo/modules/yoga"
)

type Serivce interface {
	NewGetListResponse(result []NaverBasicInfoDAO, yogaSort [][]yoga.YogaDAO) []GetListResponse
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

func (service) NewGetListResponse(result []NaverBasicInfoDAO, yogaSort [][]yoga.YogaDAO) []GetListResponse {
	response := make([]GetListResponse, 0)
	for i, v := range result {
		sort := []string{}
		for _, j := range yogaSort[i] {
			sort = append(sort, j.Name)
		}
		response = append(response, GetListResponse{
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
