package request

import "gorm.io/gorm"

type PaginationParam struct {
	Page  *int `form:"page"`
	Limit *int `form:"limit"`
}

func (p *PaginationParam) Paginate(db *gorm.DB) *gorm.DB {
	limit := 50
	offset := 0
	if p.Limit != nil {
		if *p.Limit > 100 {
			limit = 100
		} else {
			limit = *p.Limit
		}
	}
	if p.Page != nil {
		offset = (*p.Page - 1) * limit
	}

	return db.Limit(limit).Offset(offset)
}

func (p *PaginationParam) GetPage() int {
	if p.Page == nil {
		return 1
	}

	return *p.Page
}

func (p *PaginationParam) GetLimit() int {
	if p.Limit == nil {
		return 50
	}

	return *p.Limit
}
