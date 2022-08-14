package dto

type LoginResDto struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expiredAt"`
}
