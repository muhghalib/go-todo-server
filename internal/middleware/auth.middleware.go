package middleware

import (
	"fmt"
	"main/internal/domains"
	"main/internal/utils"
	"main/pkg/config"
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization, secretKey, accessToken := c.Get("Authorization"), config.Get().Jwt.Secret, ""

		if authorization == "" {
			return utils.CreateError(fiber.StatusUnauthorized, "Authorization header is missing")
		}

		fmt.Sscanf(authorization, "Bearer %s", &accessToken)

		token := utils.Token{SecretKey: secretKey}

		var claims jwt.MapClaims

		if err := token.Verify(accessToken, &claims); err != nil {
			return utils.CreateError(fiber.StatusUnauthorized, err.Error())
		}

		if ok := slices.Contains[[]string](roles, claims["role"].(string)); !ok {
			return utils.CreateError(fiber.StatusForbidden, "You are not allowed to access this route")
		}

		me := domains.Me{
			Sub:  int64(claims["sub"].(float64)),
			Role: claims["role"].(string),
		}

		c.Locals("me", me)

		return c.Next()
	}
}
