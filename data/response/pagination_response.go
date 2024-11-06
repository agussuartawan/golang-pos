package response

type PaginationResponse struct {
	Query     *string `json:"query"`
	Page      int     `json:"page"`
	Limit     int     `json:"limit"`
	Next      *int    `json:"next"`
	Prev      *int    `json:"prev"`
	SortBy    string  `json:"sortBy"`
	Sort      string  `json:"sort"`
	TotalData *int64  `json:"totalData"`
	TotalPage *int    `json:"totalPage"`
}
