package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func CreateWorkspace(c *fiber.Ctx) error {
	// uuid := fmt.Sprintf("%s", c.Locals("uuid"))

	return c.JSON(fiber.Map{"message": "success"})
}
