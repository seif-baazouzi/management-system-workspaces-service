package handlers

import (
	"fmt"
	"workspaces-service/src/models"
	"workspaces-service/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateWorkspace(c *fiber.Ctx) error {
	userID := fmt.Sprintf("%s", c.Locals("uuid"))

	// parse body
	var body models.WorkspaceBody

	err := c.BodyParser(&body)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// check values
	if body.Workspace == "" {
		return c.JSON(fiber.Map{"workspace": "Must not be empty"})
	}

	// check if parent
	if body.ParentWorkspace != uuid.Nil {
		isExit, err := models.IsWorkspaceExist(body.ParentWorkspace)

		if err != nil {
			return utils.ServerError(c, err)
		}

		if !isExit {
			return c.JSON(fiber.Map{"parentWorkspace": "This workspace does not exist"})
		}
	}

	// create workspace
	workspaceID, err := models.CreateWorkspace(body, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success", "workspaceID": workspaceID})
}
