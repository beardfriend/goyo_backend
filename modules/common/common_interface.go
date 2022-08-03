package common

type PaginationInfo struct {
	RowCount  int64 `json:"rowCount"`
	PageCount int64 `json:"pageCount"`
	Page      int64 `json:"page"`
	PageSize  int   `json:"pageSize"`
}
