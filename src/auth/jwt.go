package auth

import (
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func IsLogin(c *fiber.Ctx) error {
	claims := getJWTClaims(c)

	if claims == nil || claims["uuid"] == nil {
		return c.JSON(fiber.Map{"message": "invalid-token"})
	}

	uuid := claims["uuid"].(string)
	c.Locals("uuid", uuid)

	return c.Next()
}

func getJWTClaims(c *fiber.Ctx) jwt.MapClaims {
	jwtToken := c.Get("X-Token")
	secret := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return secret, nil
	})

	if err != nil || !token.Valid {
		return nil
	}

	return token.Claims.(jwt.MapClaims)
}
