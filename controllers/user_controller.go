package controllers

import (
	"go-project/database"
	"go-project/models"
	"go-project/utils"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []models.User

	tx := db.Find(&users)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(tx.Error.Error())
	}
	return c.Status(200).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user models.User

	id := c.Params("id")

	tx := db.Where("id = ?", id).First(&user)

	if tx.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(tx.Error.Error())
	}

	return c.Status(200).JSON(user)
}

func AddUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	errors := utils.Validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	tx := db.Create(&user)

	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(tx.Error.Error())
	}

	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user models.User

	id := c.Params("id")
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	errors := utils.Validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	tx := db.Model(models.User{}).Where("id=?", id).Updates(&user)

	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(tx.Error.Error())
	}

	return c.Status(200).JSON(user)
}

func RemoveUser(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user models.User

	result := db.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
