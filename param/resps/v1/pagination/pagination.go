package pagination

// Response 分页器响应体
type Response struct {
	CurrentPage int `json:"currentPage"`
	ItemPerPage int `json:"itemPerPage"`
	TotalPage   int `json:"totalPage"`
}