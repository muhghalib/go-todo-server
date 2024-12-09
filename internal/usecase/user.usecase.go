package usecase

import (
	"errors"
	"fmt"
	"main/internal/domains"
	"main/internal/entities"
	"main/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserUseCase struct {
	repo *gorm.DB
}

func NewUserUseCase(db *gorm.DB) domains.UserUseCase {
	return &UserUseCase{repo: db}
}

func (u *UserUseCase) FindAll(query domains.FindAllUserQuery, users *utils.Paginated[entities.User]) *fiber.Error {
	search, pagination := fmt.Sprintf("%%%s%%", query.Search), query.Pagination

	if err := u.repo.Where("name LIKE ?", search).Scopes(utils.Paginate(pagination, users)).Find(&users.Data).Error; err != nil {
		return utils.CreateError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func (u *UserUseCase) Find(id int64, user *entities.User) *fiber.Error {
	if err := u.repo.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.CreateError(fiber.StatusConflict, "user could not be found")
		} else {
			return utils.CreateError(fiber.StatusInternalServerError, err.Error())
		}
	}

	return nil
}
