package handlers

import (
	"fmt"
	"workspaces-service/src/models"
	"workspaces-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func GetSingleWorkspace(c *fiber.Ctx) error {
	uuid := fmt.Sprintf("%s", c.Locals("uuid"))
	workspaceID := c.Params("workspaceID")

	// get workspace
	workspace, exist, err := models.GetSingleWorkspace(uuid, workspaceID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !exist {
		return c.Status(404).JSON(fiber.Map{"workspaceID": "Workspace does not exist"})
	}

	return c.JSON(fiber.Map{"workspace": workspace})
}
