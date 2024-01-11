package handler

import (
	"errors"
	"net/mail"
	"time"

	"github.com/Nohty/api/config"
	"github.com/Nohty/api/database"
	"github.com/Nohty/api/model"
	"github.com/Nohty/api/utils"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func getUserByEmail(e string) (*model.User, error) {
	db := database.DB

	var user model.User
	if err := db.Where(&model.User{Email: e}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func getUserByUsername(u string) (*model.User, error) {
	db := database.DB

	var user model.User
	if err := db.Where(&model.User{Username: u}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identifier string `json:"identifier" validate:"required,ascii"`
		Password   string `json:"password" validate:"required,ascii"`
	}

	input := new(LoginInput)

	if err := utils.ParseBodyAndValidate(c, input); err != nil {
		return err
	}

	userModel, err := new(model.User), *new(error)

	if isEmail(input.Identifier) {
		userModel, err = getUserByEmail(input.Identifier)
	} else {
		userModel, err = getUserByUsername(input.Identifier)
	}

	if userModel == nil {
		return utils.Response(c, fiber.StatusUnauthorized, "User not found", err)
	}

	if !checkPasswordHash(input.Password, userModel.Password) {
		return utils.Response(c, fiber.StatusUnauthorized, "Invalid password", nil)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userModel.Username
	claims["user_id"] = userModel.ID
	claims["permission"] = userModel.Permission
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return utils.Response(c, fiber.StatusOK, "success", t)
}
