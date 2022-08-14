package helper

import (
	"gorm.io/gorm"
	"majoo/dto"
)

func Paginate(request dto.PaginationReqDto) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := request.Page
		if page == 0 {
			page = 1
		}

		pageSize := request.Limit
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
