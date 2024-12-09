package handler

import (
	"main/internal/domains"
	"main/internal/dto"
	"main/internal/usecase"
	"main/internal/utils"
	"main/pkg/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthHandler struct {
	useCase domains.AuthUseCase
}

func NewAuthHandler(router fiber.Router, db *gorm.DB) domains.AuthHandler {
	handler := &AuthHandler{useCase: usecase.NewAuthUseCase(db)}

	router.Post("/auth/login", handler.Login)
	router.Post("/auth/register", handler.Register)

	return handler
}

func (d *AuthHandler) Login(c *fiber.Ctx) error {
	authLoginDto, accessToken := dto.LoginAuthDto{}, ""

	if err := utils.Validate(c, &authLoginDto); err != nil {
		return err
	}

	if err := d.useCase.Login(authLoginDto, &accessToken); err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Path:     "/",
		Name:     "accessToken",
		Value:    accessToken,
		SameSite: "strict",
		Secure:   config.Get().App.Env == "production",
		MaxAge:   24 * 60 * 60 * 1000, /* 1 day */
	})

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message":     "Login successful",
		"accessToken": accessToken,
	})
}

func (d *AuthHandler) Register(c *fiber.Ctx) error {
	registerAuthDto := dto.RegisterAuthDto{}

	if err := utils.Validate(c, &registerAuthDto); err != nil {
		return err
	}

	if err := d.useCase.Register(registerAuthDto); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Register successful",
	})
}
