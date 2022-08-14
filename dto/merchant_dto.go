package dto

type MerchantResDto struct {
	MerchantName string  `json:"merchant_name"`
	OutletName   string  `json:"outlet_name"`
	Omzet        float64 `json:"omzet"`
}

type MerchantReqDto struct {
	MerchantId int64 `json:"merchant_id" uri:"merchantId"`
}

type MerchantsReqDto struct {
	PaginationReqDto
}
