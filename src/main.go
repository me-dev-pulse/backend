package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/me-dev-pulse/backend/src/database"
	"github.com/me-dev-pulse/backend/src/handlers"
	"github.com/me-dev-pulse/backend/src/services"
)

func main() {
	godotenv.Load()
	database.InitDB()

	services.StartMonitor(1 * time.Minute)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4321",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("DevPulse API working 🚀")
	})
	api.Get("/projects", handlers.GetProjects)
	api.Get("/projects/:id/stats", handlers.GetProjectStats)
	api.Post("/projects", handlers.CreateProjectHandler)
	api.Delete("/projects/:id", handlers.DeleteProjectHandler)
	api.Get("/projects/summary", handlers.GetProjectsSummary)

	log.Fatal(app.Listen(":3000"))
}