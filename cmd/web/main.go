package main

import (
	"errors"
	"fmt"
	"log"
	"main/internal/api"
	"main/pkg/config"
	"main/pkg/database"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func bootstrap(config *config.Config) {
	db := database.New(config.Database)

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := err.Error()

			var e *fiber.Error

			if errors.As(err, &e) {
				code = e.Code
				message = e.Message
			}

			if err != nil {
				ctx.Status(code).JSON(fiber.Map{
					"statusCode": code,
					"message":    message,
				})
			}

			return nil
		},
	})

	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))

	api.New(app, db)

	port := fmt.Sprintf(":%v", config.Server.Port)

	if err := app.Listen(port); err != nil {
		log.Printf("Failed to start server: %v", err)

		os.Exit(1)
	}
}

func main() {
	config.Load()

	cfg := config.Get()

	bootstrap(cfg)
}
