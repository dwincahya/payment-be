package models

type WebResponse[T any] struct {
	Code    int           `json:"code"`
	Data    T             `json:"data"`
	Paging  *PageMetadata `json:"paging,omitempty"`
	Message string        `json:"message"`
	Errors  string        `json:"errors,omitempty"`
}

type PageMetadata struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalItem int `json:"total_item"`
	TotalPage int `json:"total_page"`
}
