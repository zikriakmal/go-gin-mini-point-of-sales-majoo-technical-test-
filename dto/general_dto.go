package dto

import "time"

type DateFilter struct {
	DateFrom time.Time `form:"date_from,default:2022-09-01" time_format:"2006-01-02"`
	DateTo   time.Time `form:"date_to,default:2022-09-30" time_format:"2006-01-02"`
}
