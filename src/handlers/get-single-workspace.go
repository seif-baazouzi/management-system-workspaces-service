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
	workspace, err := models.GetSingleWorkspace(uuid, workspaceID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"workspace": workspace})
}
