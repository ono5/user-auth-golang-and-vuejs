// routes.go
package routes

import (
	"auth-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Home)
	app.Get("/other", controllers.Other)
}
