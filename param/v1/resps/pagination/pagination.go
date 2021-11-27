package pagination

// Response 分页器响应体
type Response struct {
	CurrentPage int `json:"currentPage"`
	ItemCounter int `json:"itemCounter"`
	TotalPage   int `json:"totalPage"`
}