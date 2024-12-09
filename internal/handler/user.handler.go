package handler

import (
	"main/internal/domains"
	"main/internal/entities"
	"main/internal/middleware"
	"main/internal/usecase"
	"main/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	useCase domains.UserUseCase
}

func NewUserHandler(router fiber.Router, db *gorm.DB) domains.UserHandler {
	handler := &UserHandler{useCase: usecase.NewUserUseCase(db)}

	router.Get("/users", middleware.Auth("admin"), handler.FindAll)
	router.Get("/users/:id", middleware.Auth("admin"), handler.Find)

	return handler
}

func (d *UserHandler) FindAll(c *fiber.Ctx) error {
	users := utils.Paginated[entities.User]{}

	query := domains.FindAllUserQuery{
		Pagination: utils.Pagination{Page: c.QueryInt("page", 1), Limit: c.QueryInt("limit", 10)},
		Search:     c.Query("search", ""),
	}

	if err := d.useCase.FindAll(query, &users); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (d *UserHandler) Find(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	user := entities.User{}

	if err != nil {
		return utils.CreateError(fiber.StatusInternalServerError, err.Error())
	}

	if err := d.useCase.Find(int64(id), &user); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
