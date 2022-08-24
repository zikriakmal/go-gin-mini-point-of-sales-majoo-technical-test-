package dto

import "time"

type MerchantReportReqDto struct {
	MerchantReqDto
	DateFilter
	PaginationReqDto
}

type MerchantReportResDto struct {
	MerchantName string    `json:"merchant_name"`
	Omzet        float64   `json:"omzet"`
	Date         time.Time `json:"date"`
}
