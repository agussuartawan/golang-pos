package request

import (
	"gorm.io/gorm"
)

type PaginationParam struct {
	Page      *int        `form:"page"`
	Limit     *int        `form:"limit"`
	TotalData *int64      `form:"-"`
	TotalPage *int        `form:"-"`
	NextPage  *int        `form:"-"`
	PrevPage  *int        `form:"-"`
	Data      interface{} `form:"-"`
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

	// count total data
	var total int64
	db.Count(&total)
	p.TotalData = &total

	// Hitung total halaman
	totalPage := int(total) / limit
	if int(total)%limit != 0 {
		totalPage++ // Tambahkan 1 jika ada sisa
	}
	p.TotalPage = &totalPage

	// Hitung halaman berikutnya dan sebelumnya
	if p.Page != nil {
		currentPage := *p.Page
		if currentPage < totalPage {
			nextPage := currentPage + 1
			p.NextPage = &nextPage
		}
		if currentPage > 1 {
			prevPage := currentPage - 1
			p.PrevPage = &prevPage
		}
	} else {
		// Default ke halaman pertama jika Page tidak diset
		firstPage := 1
		p.Page = &firstPage
		p.NextPage = &firstPage
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

func (p *PaginationParam) SetData(data interface{}) PaginationParam {
	p.Data = data
	return *p
}
