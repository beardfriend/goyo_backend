package common

import "sync"

type Service interface {
	NewPaginationInfo(resultLength int, total int64, pageNo int, rowCount int) PaginationInfo
}

type service struct{}

func (service) NewPaginationInfo(
	resultLength int,
	total int64,
	pageNo int,
	rowCount int,
) PaginationInfo {
	pageCount := total/int64(rowCount) + 1
	if total%int64(rowCount) == 0 {
		pageCount = total / int64(rowCount)
	}
	pageInfo := PaginationInfo{
		PageSize:  resultLength,
		PageCount: pageCount,
		Page:      int64(pageNo),
		RowCount:  total,
	}

	return pageInfo
}

// ------------------- SingleTon -------------------

var (
	serviceInstance Service
	serviceOnce     sync.Once
)

func GetService() Service {
	if serviceInstance != nil {
		return serviceInstance
	}
	serviceOnce.Do(func() {
		serviceInstance = &service{}
	})
	return serviceInstance
}
