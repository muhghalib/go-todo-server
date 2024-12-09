package handler

import (
	"main/internal/domains"
	"main/internal/dto"
	"main/internal/entities"
	"main/internal/middleware"
	"main/internal/usecase"
	"main/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TaskHandler struct {
	useCase domains.TaskUseCase
}

func NewTaskHandler(router fiber.Router, db *gorm.DB) domains.TaskHandler {
	handler := &TaskHandler{useCase: usecase.NewTaskUseCase(db)}

	router.Get("/tasks", middleware.Auth("default"), handler.FindAll)
	router.Get("/tasks/:id", middleware.Auth("default"), handler.Find)

	router.Post("/tasks", middleware.Auth("default"), handler.Create)
	router.Patch("/tasks/:id", middleware.Auth("default"), handler.Update)

	return handler
}

func (t *TaskHandler) FindAll(c *fiber.Ctx) error {
	me := c.Locals("me").(domains.Me)

	tasks := []entities.Task{}

	if err := t.useCase.FindAll(me.Sub, &tasks); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(tasks)
}

func (t *TaskHandler) Find(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return utils.CreateError(fiber.StatusInternalServerError, err.Error())
	}

	me := c.Locals("me").(domains.Me)

	task := entities.Task{}

	if err := t.useCase.Find(me.Sub, int64(id), &task); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

func (t *TaskHandler) Create(c *fiber.Ctx) error {
	me := c.Locals("me").(domains.Me)

	createTaskDto := dto.CreateTaskDto{}

	if err := utils.Validate(c, &createTaskDto); err != nil {
		return err
	}

	if err := t.useCase.Create(me.Sub, createTaskDto); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).SendString("created")
}

func (t *TaskHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return utils.CreateError(fiber.StatusInternalServerError, err.Error())
	}

	me := c.Locals("me").(domains.Me)

	updateTaskDto := dto.UpdateTaskDto{}

	if err := utils.Validate(c, &updateTaskDto); err != nil {
		return err
	}

	if err := t.useCase.Update(me.Sub, int64(id), updateTaskDto); err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).SendString("updated")

}
