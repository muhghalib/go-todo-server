package usecase

import (
	"errors"
	"time"

	"main/internal/domains"
	"main/internal/dto"
	"main/internal/entities"
	"main/internal/utils"
	"main/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthUseCase struct {
	repo *gorm.DB
}

func NewAuthUseCase(db *gorm.DB) domains.AuthUseCase {
	return &AuthUseCase{repo: db}
}

func hashPassword(plain *string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(*plain), bcrypt.DefaultCost)

	*plain = string(hash)

	return err
}

func comparePassword(plain string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}

func createAccessToken(user entities.User, accessToken *string) error {
	secretKey := config.Get().Jwt.Secret

	token := utils.Token{SecretKey: secretKey, Claims: jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}}

	if err := token.Generate(accessToken); err != nil {
		return err
	}

	return nil
}

func (a *AuthUseCase) Login(loginAuthDto dto.LoginAuthDto, accessToken *string) *fiber.Error {
	user := entities.User{}

	if err := a.repo.Where("email = ?", loginAuthDto.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.CreateError(fiber.StatusConflict, "user email could not be found")
		} else {
			return utils.CreateError(fiber.StatusInternalServerError, err.Error())
		}
	}

	if err := comparePassword(loginAuthDto.Password, user.Password); err != nil {
		return utils.CreateError(fiber.StatusUnauthorized, "invalid user password")
	}

	if err := createAccessToken(user, accessToken); err != nil {
		return utils.CreateError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func (a *AuthUseCase) Register(registerAuthDto dto.RegisterAuthDto) *fiber.Error {
	user := entities.User{
		Name:     registerAuthDto.Name,
		Email:    registerAuthDto.Email,
		Password: registerAuthDto.Password,
	}

	if err := a.repo.Where("email = ?", user.Email).First(&user).Error; err == nil {
		return utils.CreateError(fiber.StatusConflict, "user email is already exist")
	}

	if err := hashPassword(&user.Password); err != nil {
		return utils.CreateError(fiber.StatusInternalServerError, err.Error())
	}

	if err := a.repo.Create(&user).Error; err != nil {
		return utils.CreateError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
