package domains

import (
	"main/internal/dto"

	"github.com/gofiber/fiber/v2"
)

type AuthUseCase interface {
	Login(dto.LoginAuthDto, *string) *fiber.Error
	Register(dto.RegisterAuthDto) *fiber.Error
}

type AuthHandler interface {
	Login(*fiber.Ctx) error
	Register(*fiber.Ctx) error
}
