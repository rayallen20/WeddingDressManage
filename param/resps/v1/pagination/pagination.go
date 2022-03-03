package pagination

import "math"

// Response 分页器响应体
type Response struct {
	CurrentPage int   `json:"currentPage"`
	ItemPerPage int   `json:"itemPerPage"`
	TotalPage   int   `json:"totalPage"`
	TotalItem   int64 `json:"totalItem"`
}

func CalcTotalPage(count int64, itemPerPage int) int {
	return int(math.Ceil(float64(count) / float64(itemPerPage)))
}
