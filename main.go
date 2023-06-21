package main

import (
	"restgo/db"
	"restgo/handlers"
	"restgo/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	db.DBConection()
	db.DB.AutoMigrate(models.Tasks{})

	app.Use(cors.New(cors.ConfigDefault))
	app.Get("/tasks", handlers.Tasks)
	app.Post("/task", handlers.New)
	app.Get("/task/:id", handlers.Task)
	app.Patch("/task/:id", handlers.Done)
	app.Put("/task/:id", handlers.Edit)
	app.Delete("/task/:id", handlers.Remove)

	app.Listen(":3000")
}
