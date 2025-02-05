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

////////////////////////////////////// request structs

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
	Task       *entity.TaskDTO `json:"Tasks"`
	StatusCode int
}

func (r *TaskGetOneResponse) GetStatus() int {
	return r.StatusCode
}

func (r *TaskGetOneResponse) GetBody() interface{} {
	return r.Task
}

type TaskGetAllResponse struct {
	Tasks      []*entity.TaskDTO `json:"task"`
	StatusCode int
}

func (r *TaskGetAllResponse) GetStatus() int {
	return r.StatusCode
}

func (r *TaskGetAllResponse) GetBody() interface{} {
	return r.Tasks
}

type TaskPostResponse struct {
	TaskPost   TaskPost
	StatusCode int
}

func (r *TaskPostResponse) GetStatus() int {
	return r.StatusCode
}

func (r *TaskPostResponse) GetBody() interface{} {
	return r.TaskPost
}

type TaskPutResponse struct {
	TaskPut    TaskPut
	StatusCode int
}

func (r *TaskPutResponse) GetStatus() int {
	return r.StatusCode
}

func (r *TaskPutResponse) GetBody() interface{} {
	return r.TaskPut
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

type TaskPost struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskPut struct {
	Id        uint      `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
}

////////////////////////////////////////////////
