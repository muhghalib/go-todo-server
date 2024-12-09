package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Validate(c *fiber.Ctx, model interface{}) *fiber.Error {
	if err := c.BodyParser(&model); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	validate := validator.New()

	if err := validate.Struct(model); err != nil {

		formattedErrors := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			formattedErrors[strings.ToLower(err.Field())] = fmt.Sprintf("failed validation with tag '%s'", err.Tag())
		}

		errorMessage, err := json.Marshal(formattedErrors)

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return fiber.NewError(fiber.StatusBadRequest, string(errorMessage))
	}

	return nil
}
