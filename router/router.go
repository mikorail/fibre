package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mikorail/fibre/handler"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/request")
	// routes
	v1.Get("/", handler.GetAllReqs)
	v1.Get("/:id", handler.GetSingleReq)
	v1.Post("/", handler.CreateReqs)
	v1.Delete("/:id", handler.DeleteReq)
}
