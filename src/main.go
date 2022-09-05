package main

import (
	"fmt"
	"os"
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

	app := fiber.New()
	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/api/v1/workspaces", auth.IsLogin, handlers.GetWorkspaces)
	app.Post("/api/v1/workspaces", auth.IsLogin, handlers.CreateWorkspace)
	app.Put("/api/v1/workspaces/:workspaceID", auth.IsLogin, handlers.UpdateWorkspace)
	app.Delete("/api/v1/workspaces/:workspaceID", auth.IsLogin, handlers.DeleteWorkspace)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
