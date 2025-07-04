package models

type WebResponse[T any] struct {
	Data   T             `json:"data"`
	Paging *PageMetadata `json:"paging,omitempty"`
	Errors string        `json:"errors,omitempty"`
}

type PageMetadata struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalItem int `json:"total_item"`
	TotalPage int `json:"total_page"`
}
