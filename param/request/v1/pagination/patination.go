package pagination

type Pagination struct {
	CurrentPage int `form:"currentPage" binding:"gt=0,required" errField:"currentPage"`
	ItemPerPage int `form:"itemPerPage" binding:"gt=0,required" errField:"itemPerPage"`
}
