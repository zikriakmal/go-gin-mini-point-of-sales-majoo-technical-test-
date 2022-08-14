package dto

type OutletResDto struct {
	ID           int64   `json:"id"`
	MerchantName string  `json:"merchant_name"`
	OutletName   string  `json:"outlet_name"`
	Omzet        float64 `json:"omzet"`
}

type OutletReqDto struct {
	MerchantReqDto
	OutletId int64 `json:"outlet_id" uri:"outletId"`
}

type OutletsReqDto struct {
	MerchantReqDto
	PaginationReqDto
	DateFilter
}
