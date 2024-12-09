package dto

import "main/internal/entities"

type CreateTaskDto struct {
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required"`
}

type UpdateTaskDto struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Status      entities.TaskStatus `json:"status"`
}
