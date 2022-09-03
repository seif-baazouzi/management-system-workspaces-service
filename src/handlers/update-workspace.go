package handlers

import (
	"fmt"
	"workspaces-service/src/models"
	"workspaces-service/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UpdateWorkspace(c *fiber.Ctx) error {
	userID := fmt.Sprintf("%s", c.Locals("uuid"))

	workspaceID := c.Params("workspaceID", uuid.Nil.String())

	// parse body
	var body models.WorkspaceBody

	err := c.BodyParser(&body)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// check if workspace is not exist
	isExit, err := models.IsWorkspaceExist(workspaceID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !isExit {
		return c.JSON(fiber.Map{"workspaceID": "This workspace does not exist"})
	}

	// check values
	if body.Workspace == "" {
		return c.JSON(fiber.Map{"workspace": "Must not be empty"})
	}

	// check if parent
	if body.ParentWorkspace != uuid.Nil {
		isExit, err := models.IsWorkspaceExist(body.ParentWorkspace.String())

		if err != nil {
			return utils.ServerError(c, err)
		}

		if !isExit {
			return c.JSON(fiber.Map{"parentWorkspace": "This workspace does not exist"})
		}
	}

	// update workspace
	err = models.UpdateWorkspace(body, workspaceID, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}
