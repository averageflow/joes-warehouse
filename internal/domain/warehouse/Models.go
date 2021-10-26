package warehouse

type PaginationInQueryParams struct {
	Limit  int64 `form:"limit"`
	Offset int64 `form:"offset"`
}
