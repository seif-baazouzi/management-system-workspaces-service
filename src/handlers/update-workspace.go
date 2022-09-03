package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func UpdateWorkspace(c *fiber.Ctx) error {
	// userID := fmt.Sprintf("%s", c.Locals("uuid"))

	return c.JSON(fiber.Map{"message": "success"})
}
