package model

type Model interface {
	GetOne(id int) *ApiResponse
	GetAll(limit int, offset int) *ApiResponse
	Post(req PostRequest) *ApiResponse
	Put(req PutRequest) *ApiResponse
	Delete(id int) *ApiResponse
}

type PostRequest interface{}
type PutRequest interface{}

type ApiResponse struct {
	JsonData   string
	StatusCode int
}
