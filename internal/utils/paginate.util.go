package utils

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Page  int
	Limit int
}

type Paginated[T interface{}] struct {
	Data        []T  `json:"data"`
	Page        int  `json:"page"`
	Limit       int  `json:"limit"`
	DataCount   int  `json:"dataCount"`
	PageCount   int  `json:"pageCount"`
	CanPrevPage bool `json:"canPrevPage"`
	CanNextPage bool `json:"canNextPage"`
}

func Paginate[T interface{}](pagination Pagination, paginated *Paginated[T]) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		page, limit := pagination.Page, pagination.Limit

		offset := (page - 1) * limit

		var dataCount int64

		if err := db.Session(&gorm.Session{SkipHooks: true}).Count(&dataCount).Session(&gorm.Session{SkipHooks: true}).Error; err != nil {
			panic("Failed to count rows: " + err.Error())
		}

		paginated.Page = page
		paginated.Limit = limit
		paginated.DataCount = int(dataCount)
		paginated.PageCount = int((dataCount + int64(limit) - 1) / int64(limit))
		paginated.CanPrevPage = page > 1
		paginated.CanNextPage = int64(page*limit) < dataCount

		return db.Offset(offset).Limit(limit).Session(&gorm.Session{SkipHooks: true})
	}
}
