package model

type Model interface {
	GetOne() *ApiResponse
	GetAll() *ApiResponse
	Post() *ApiResponse
	Put() *ApiResponse
	Delete() *ApiResponse
}

type ApiResponse struct {
	JsonData   string
	StatusCode int
}
