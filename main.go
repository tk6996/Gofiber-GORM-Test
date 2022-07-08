package main

import (
	"go-project/database"
	"go-project/routers"
	"go-project/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	utils.Validate = validator.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")
	routers.UserRouter(v1)
	database.InitDatabase()
	app.Listen(":3000")
}
