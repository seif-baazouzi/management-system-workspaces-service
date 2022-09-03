package handlers

import (
	"fmt"
	"workspaces-service/src/models"
	"workspaces-service/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DeleteWorkspace(c *fiber.Ctx) error {
	userID := fmt.Sprintf("%s", c.Locals("uuid"))
	workspaceID := c.Params("workspaceID", uuid.Nil.String())

	// check if workspace is not exist
	isExit, err := models.IsWorkspaceExist(workspaceID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !isExit {
		return c.JSON(fiber.Map{"message": "This workspace does not exist"})
	}

	// delete workspace
	err = models.DeleteWorkspace(workspaceID, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}
