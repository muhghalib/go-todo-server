package usecase

import (
	"errors"
	"fmt"
	"main/internal/domains"
	"main/internal/dto"
	"main/internal/entities"
	"main/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TaskUseCase struct {
	repo *gorm.DB
}

func NewTaskUseCase(db *gorm.DB) domains.TaskUseCase {
	return &TaskUseCase{repo: db}
}

func (t *TaskUseCase) FindAll(userId int64, tasks *[]entities.Task) *fiber.Error {
	if err := t.repo.Where("user_id = ?", userId).Omit("User").Find(tasks).Error; err != nil {
		return utils.CreateError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func (t *TaskUseCase) Find(userId int64, taskId int64, task *entities.Task) *fiber.Error {
	if err := t.repo.Where("user_id = ? & id = ?", userId, taskId).First(task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.CreateError(fiber.StatusNotFound, "task could not be found")
		} else {
			return utils.CreateError(fiber.StatusInternalServerError, err.Error())
		}
	}

	return nil
}

func (t *TaskUseCase) Create(userId int64, createTaskDto dto.CreateTaskDto) *fiber.Error {
	task := entities.Task{
		Title:       createTaskDto.Title,
		Description: createTaskDto.Description,
		UserID:      uint(userId),
		DueDate:     time.Now(),
	}

	if err := t.repo.Create(&task).Error; err != nil {
		return utils.CreateError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func (t *TaskUseCase) Update(userId int64, taskId int64, updateTaskDto dto.UpdateTaskDto) *fiber.Error {
	task := entities.Task{}

	if err := t.Find(userId, taskId, &task); err != nil {
		return err
	}

	fmt.Println(updateTaskDto.Description)

	task.Title = updateTaskDto.Title
	task.Description = updateTaskDto.Description
	task.Status = updateTaskDto.Status

	if err := t.repo.Updates(&task).Error; err != nil {
		return utils.CreateError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
