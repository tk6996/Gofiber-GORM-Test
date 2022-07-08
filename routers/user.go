package routers

import (
	"go-project/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func UserRouter(app fiber.Router) {
	auth := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "772565",
		},
	})
	app.Get("/users", controllers.GetUsers)
	app.Get("/user/:id", controllers.GetUser)
	app.Post("/user", auth, controllers.AddUser)
	app.Put("/user/:id", auth, controllers.UpdateUser)
	app.Delete("/user/:id", auth, controllers.RemoveUser)
}
