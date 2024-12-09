package api

import (
	"main/internal/handler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func New(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api/v1")

	handler.NewAuthHandler(api, db)
	handler.NewUserHandler(api, db)
	handler.NewTaskHandler(api, db)
}
