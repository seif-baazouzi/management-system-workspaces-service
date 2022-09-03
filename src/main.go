package main

import (
	"workspaces-service/src/auth"
	"workspaces-service/src/db"
	"workspaces-service/src/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db.InitDataBase()
	defer db.CloseDataBase()
	db.Migrations()

	app := fiber.New()
	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/api/v1/workspaces", auth.IsLogin, handlers.GetWorkspaces)
	app.Post("/api/v1/workspaces", auth.IsLogin, handlers.CreateWorkspace)

	app.Listen(":3000")
}
