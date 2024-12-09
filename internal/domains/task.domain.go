package domains

import (
	"main/internal/dto"
	"main/internal/entities"

	"github.com/gofiber/fiber/v2"
)

type TaskUseCase interface {
	FindAll(int64, *[]entities.Task) *fiber.Error
	Find(int64, int64, *entities.Task) *fiber.Error
	Create(int64, dto.CreateTaskDto) *fiber.Error
	Update(int64, int64, dto.UpdateTaskDto) *fiber.Error
}

type TaskHandler interface {
	FindAll(*fiber.Ctx) error
	Find(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	Update(*fiber.Ctx) error
}
