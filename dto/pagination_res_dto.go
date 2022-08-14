package dto

type PaginationRes[T interface{}] struct {
	MetaData `json:"_metadata"`
	Records  []T `json:"records"`
}
type MetaData struct {
	Page    int   `json:"page"`
	PerPage int   `json:"per_page"`
	Total   int64 `json:"total"`
}
