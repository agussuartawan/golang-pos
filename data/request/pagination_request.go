package request

import (
	"gorm.io/gorm"
)

type PaginationParam struct {
	Page      int         `form:"page"`
	Limit     int         `form:"limit"`
	TotalData *int64      `form:"-"`
	TotalPage *int        `form:"-"`
	NextPage  *int        `form:"-"`
	PrevPage  *int        `form:"-"`
	Data      interface{} `form:"-"`
}

func (p *PaginationParam) Paginate(db *gorm.DB) *gorm.DB {
	limit := 50
	offset := 0
	if p.Limit > 0 {
		if p.Limit > 100 {
			limit = 100
		} else {
			limit = p.Limit
		}
	}

	if p.Page > 0 {
		offset = (p.Page - 1) * limit
	} else {
		// Default ke halaman pertama jika Page tidak diset
		p.Page = 1
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
	currentPage := p.Page
	if currentPage < totalPage {
		nextPage := currentPage + 1
		p.NextPage = &nextPage
	}
	if currentPage > 1 {
		prevPage := currentPage - 1
		p.PrevPage = &prevPage
	}

	return db.Limit(limit).Offset(offset)
}

func (p *PaginationParam) SetData(data interface{}) PaginationParam {
	p.Data = data
	return *p
}
