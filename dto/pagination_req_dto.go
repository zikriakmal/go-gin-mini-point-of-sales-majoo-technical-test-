package dto

type PaginationReqDto struct {
	Page  int `json:"page" form:"page,default:1"`
	Limit int `json:"limit" form:"limit,default:1"`
}
