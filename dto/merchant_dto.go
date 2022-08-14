package dto

type MerchantResDto struct {
	ID           int64   `json:"id"`
	MerchantName string  `json:"merchant_name"`
	Omzet        float64 `json:"omzet"`
}

type MerchantReqDto struct {
	MerchantId int64 `json:"merchant_id" uri:"merchantId"`
}

type MerchantsReqDto struct {
	PaginationReqDto
}
