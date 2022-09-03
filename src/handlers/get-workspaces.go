package handlers

import (
	"fmt"
	"workspaces-service/src/models"
	"workspaces-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func GetWorkspaces(c *fiber.Ctx) error {
	uuid := fmt.Sprintf("%s", c.Locals("uuid"))

	// get user
	workspaces, err := models.GetWorkspaces(uuid)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"workspaces": workspaces})
}
