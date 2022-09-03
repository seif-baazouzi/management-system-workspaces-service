package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func DeleteWorkspace(c *fiber.Ctx) error {
	// userID := fmt.Sprintf("%s", c.Locals("uuid"))
	// workspaceID := c.Params("workspaceID", uuid.Nil.String())

	return c.JSON(fiber.Map{"message": "success"})
}
