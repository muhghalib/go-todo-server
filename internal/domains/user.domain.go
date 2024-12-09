package domains

import (
	"main/internal/entities"
	"main/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type FindAllUserQuery struct {
	utils.Pagination
	Search string
}

type UserUseCase interface {
	FindAll(FindAllUserQuery, *utils.Paginated[entities.User]) *fiber.Error
	Find(int64, *entities.User) *fiber.Error
}

type UserHandler interface {
	FindAll(*fiber.Ctx) error
	Find(*fiber.Ctx) error
}
