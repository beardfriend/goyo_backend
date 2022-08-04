package academy

import (
	"sync"

	"goyo/modules/common"
)

type Serivce interface {
	NewGetListResponse(result []NaverBasicInfoDTO, total int64, query *GetListQuery) GetListResponse
}

type service struct{}

func (service) NewGetListResponse(result []NaverBasicInfoDTO, total int64, query *GetListQuery) GetListResponse {
	var response GetListResponse

	pageInfo := common.PaginationInfo{
		PageSize:  len(result),
		PageCount: total/int64(query.RowCount) + 1,
		Page:      int64(query.PageNo),
		RowCount:  total,
	}

	response.List = result
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
