package model

type Model interface {
	GetOne(id int) (*ApiResponse, error)
	GetAll(limit int, offset int) (*ApiResponse, error)
	Post(req PostRequest) (*ApiResponse, error)
	Put(req PutRequest) (*ApiResponse, error)
	Delete(req DeleteRequest) (*ApiResponse, error)
}

type PostRequest interface{}
type PutRequest interface{}
type DeleteRequest interface{}

type ApiResponse struct {
	JsonData   string
	StatusCode int
}
