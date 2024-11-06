package request

import (
	"github.com/agussuartawan/golang-pos/data/response"
	"gorm.io/gorm"
)

type PaginationParam struct {
	Query     *string `form:"query"`
	Page      int     `form:"page"`
	Limit     int     `form:"limit"`
	SortBy    string  `form:"sortBy"`
	Sort      string  `form:"sort"`
	TotalData *int64  `form:"-"`
	TotalPage *int    `form:"-"`
	NextPage  *int    `form:"-"`
	PrevPage  *int    `form:"-"`
}

func (p *PaginationParam) ToResponse() *response.PaginationResponse {
	return &response.PaginationResponse{
		Query:     p.Query,
		Page:      p.Page,
		Limit:     p.Limit,
		Next:      p.NextPage,
		Prev:      p.PrevPage,
		SortBy:    p.SortBy,
		Sort:      p.Sort,
		TotalPage: p.TotalPage,
		TotalData: p.TotalData,
	}
}

func (p *PaginationParam) Paginate(db *gorm.DB) *gorm.DB {
	limit := 20
	offset := 0
	if p.Limit > 0 {
		if p.Limit > 100 {
			limit = 100
		} else {
			limit = p.Limit
		}
	} else {
		p.Limit = limit
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

	if p.SortBy == "" {
		p.SortBy = "createdAt"
	}
	if p.Sort == "" || p.Sort == "desc" {
		p.Sort = "desc"
	} else {
		p.Sort = "asc"
	}

	return db.Limit(limit).Offset(offset)
}
