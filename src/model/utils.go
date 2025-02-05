package model

import (
	"repos/task_manager/src/entity"
)

// CreateDTOFromTaskModel creates custom entity.TaskDTO struct from gorm.Model struct
func CreateDTOFromTaskModel(task *entity.Task) *entity.TaskDTO {
	var dto entity.TaskDTO
	dto.Id = task.ID
	dto.Title = task.Title
	dto.Description = task.Description
	dto.CreatedAt = task.CreatedAt
	dto.UpdatedAt = task.UpdatedAt
	return &dto
}
