package di

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ratheeshkumar25/routing_featureMT5VsHS/pkg/handlers"
)

func Init() {
	app := fiber.New()

	app.Post("/api/v1/rules", handlers.AddRule)
	app.Put("/api/v1/rules/:id", handlers.UpdateRule)
	app.Delete("/api/v1/rules/:id", handlers.DeleteRule)
	app.Get("/api/v1/rules", handlers.GetRules)

	app.Listen(":8080")
}
