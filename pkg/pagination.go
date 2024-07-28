package pkg

import (
	"gorm.io/gorm"
)

func Paginate(offset, limit int, sort string, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit).Order(sort)
	}
}

type Pagination struct {
	Offset  int
	PerPage int
	Page    int
}

func PaginationBuilder(perPage, page int) *Pagination {
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * perPage

	paginator := Pagination{
		Offset:  offset,
		PerPage: perPage,
		Page:    page,
	}
	return &paginator
}
