package handler

import (
	"github.com/Nohty/api/database"
	"github.com/Nohty/api/model"
	"github.com/Nohty/api/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	userId, _ := utils.GetUserID(c)
	permission := utils.GetPermissionFromDB(userId)

	if (uint(id) != userId && !utils.HasPermission(permission, utils.IsAdmin)) || permission == 0 {
		return fiber.ErrForbidden
	}

	db := database.DB

	var user model.User
	db.Preload("Contacts").Preload("Address").Find(&user, id)

	if user.Username == "" {
		return utils.Response(c, fiber.StatusNotFound, "User not found", nil)
	}

	return utils.Response(c, fiber.StatusOK, "success", user)
}

func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Username   string `json:"username" validate:"required,ascii"`
		Email      string `json:"email" validate:"required,email"`
		Phone      string `json:"phone" validate:"required,e164"`
		Password   string `json:"password" validate:"required,ascii,min=6"`
		WalletAddr string `json:"wallet_addr" validate:"required,ascii"`
		Street     string `json:"street" validate:"required,ascii"`
		City       string `json:"city" validate:"required,ascii"`
		Postcode   string `json:"postcode" validate:"required,ascii"`
		Number     string `json:"number" validate:"required,ascii"`
	}

	input := new(NewUser)

	if err := utils.ParseBodyAndValidate(c, input); err != nil {
		return err
	}

	user := model.User{
		Username:   input.Username,
		Email:      input.Email,
		Phone:      input.Phone,
		WalletAddr: input.WalletAddr,
		Address: model.Address{
			Street:   input.Street,
			City:     input.City,
			Postcode: input.Postcode,
			Number:   input.Number,
		},
		Contacts: []model.User{},
	}

	hash, err := hashPassword(input.Password)
	if err != nil {
		return utils.Response(c, fiber.StatusInternalServerError, "Error hashing password", nil)
	}

	user.Password = hash
	user.Permission = utils.NewPermissionFlags(utils.IsUser)

	db := database.DB
	if err := db.Create(&user).Error; err != nil {
		return utils.Response(c, fiber.StatusInternalServerError, "Error creating user", err)
	}

	return utils.Response(c, fiber.StatusCreated, "User successfully created", user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	type UpdateUser struct {
		Username   string `json:"username" validate:"required,ascii"`
		Email      string `json:"email" validate:"required,email"`
		Phone      string `json:"phone" validate:"required,e164"`
		WalletAddr string `json:"wallet_addr" validate:"required,ascii"`
		Street     string `json:"street" validate:"required,ascii"`
		City       string `json:"city" validate:"required,ascii"`
		Postcode   string `json:"postcode" validate:"required,ascii"`
		Number     string `json:"number" validate:"required,ascii"`
	}

	input := new(UpdateUser)

	if err := utils.ParseBodyAndValidate(c, input); err != nil {
		return err
	}

	userId, _ := utils.GetUserID(c)
	permission := utils.GetPermissionFromDB(userId)

	if (uint(id) != userId && !utils.HasPermission(permission, utils.IsAdmin)) || permission == 0 {
		return fiber.ErrForbidden
	}

	db := database.DB

	var user model.User
	if err := db.Preload("Contacts").Preload("Address").Find(&user, id).Error; err != nil {
		return utils.Response(c, fiber.StatusNotFound, "User not found", nil)
	}

	user.Username = input.Username
	user.Email = input.Email
	user.Phone = input.Phone
	user.WalletAddr = input.WalletAddr
	user.Address.Street = input.Street
	user.Address.City = input.City
	user.Address.Postcode = input.Postcode
	user.Address.Number = input.Number

	if err := db.Save(&user).Error; err != nil {
		return utils.Response(c, fiber.StatusInternalServerError, "Error updating user", err)
	}

	return utils.Response(c, fiber.StatusOK, "User successfully updated", user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	type DeleteUser struct {
		Password string `json:"password" validate:"required,ascii,min=6"`
	}

	input := new(DeleteUser)

	if err := utils.ParseBodyAndValidate(c, input); err != nil {
		return err
	}

	userId, _ := utils.GetUserID(c)
	permission := utils.GetPermissionFromDB(userId)

	if (uint(id) != userId && !utils.HasPermission(permission, utils.IsAdmin)) || permission == 0 {
		return fiber.ErrForbidden
	}

	db := database.DB

	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		return utils.Response(c, fiber.StatusNotFound, "User not found", nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return utils.Response(c, fiber.StatusUnauthorized, "Invalid password", nil)
	}

	if err := db.Delete(&user).Error; err != nil {
		return utils.Response(c, fiber.StatusInternalServerError, "Error deleting user", err)
	}

	return utils.Response(c, fiber.StatusOK, "User successfully deleted", nil)
}

func UpdatePassword(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	type UpdatePassword struct {
		OldPassword string `json:"old_password" validate:"required,ascii,min=6"`
		NewPassword string `json:"new_password" validate:"required,ascii,min=6"`
	}

	input := new(UpdatePassword)

	if err := utils.ParseBodyAndValidate(c, input); err != nil {
		return err
	}

	userId, _ := utils.GetUserID(c)
	permission := utils.GetPermissionFromDB(userId)

	if (uint(id) != userId && !utils.HasPermission(permission, utils.IsAdmin)) || permission == 0 {
		return fiber.ErrForbidden
	}

	db := database.DB

	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		return utils.Response(c, fiber.StatusNotFound, "User not found", nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword)); err != nil {
		return utils.Response(c, fiber.StatusUnauthorized, "Invalid password", nil)
	}

	hash, err := hashPassword(input.NewPassword)
	if err != nil {
		return utils.Response(c, fiber.StatusInternalServerError, "Error hashing password", nil)
	}

	user.Password = hash

	if err := db.Save(&user).Error; err != nil {
		return utils.Response(c, fiber.StatusInternalServerError, "Error updating password", err)
	}

	return utils.Response(c, fiber.StatusOK, "Password successfully updated", nil)
}
