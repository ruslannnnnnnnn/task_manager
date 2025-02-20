package model

import (
	"repos/task_manager/src/entity"
	"time"
)

type Model interface {
	GetOne(id int) (ApiResponse, error)
	GetAll(limit int, offset int) (ApiResponse, error)
	Post(req PostRequest) (ApiResponse, error)
	Put(req PutRequest) (ApiResponse, error)
	Delete(req DeleteRequest) (ApiResponse, error)
}

type PostRequest interface{}
type PutRequest interface{}
type DeleteRequest interface{}

type ApiResponse interface {
	GetStatus() int
	GetBody() interface{}
}

////////////////////////////////////// request body structs

type TaskPostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskPutRequest struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskDeleteRequest struct {
	Ids []int `json:"ids"`
}

/////////////////////////////////// response structs

type TaskGetOneResponse struct {
	Body       *entity.TaskDTO `json:"tasks"`
	StatusCode int
}

func (r *TaskGetOneResponse) GetStatus() int {
	return r.StatusCode
}

func (r *TaskGetOneResponse) GetBody() interface{} {
	return r.Body
}

type TaskGetAllResponse struct {
	Body       []*entity.TaskDTO `json:"task"`
	StatusCode int
}

func (r *TaskGetAllResponse) GetStatus() int {
	return r.StatusCode
}

func (r *TaskGetAllResponse) GetBody() interface{} {
	return r.Body
}

type TaskPostResponse struct {
	Body       TaskPostResponseBody
	StatusCode int
}

func (r *TaskPostResponse) GetStatus() int {
	return r.StatusCode
}

func (r *TaskPostResponse) GetBody() interface{} {
	return r.Body
}

type TaskPutResponse struct {
	Body       TaskPutResponseBody
	StatusCode int
}

func (r *TaskPutResponse) GetStatus() int {
	return r.StatusCode
}

func (r *TaskPutResponse) GetBody() interface{} {
	return r.Body
}

type TaskDeleteResponse struct {
	StatusCode int
}

func (r *TaskDeleteResponse) GetStatus() int {
	return r.StatusCode
}

func (r *TaskDeleteResponse) GetBody() interface{} {
	return nil
}

type TaskPostResponseBody struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskPutResponseBody struct {
	Id        uint      `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
}

////////////////////////////////////////////////
