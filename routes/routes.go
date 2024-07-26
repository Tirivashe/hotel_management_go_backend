package routes

import "github.com/gofiber/fiber/v2"

func LoadRoutes(app *fiber.App) {
	api := app.Group("api/v1")
	api.Get("/", healthCheck)
}

func healthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Welcome to Hotel Management Backend"})
}
