package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"workspaces-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func IsLogin(c *fiber.Ctx) error {
	conn := http.Client{Timeout: time.Duration(3) * time.Second}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/user", os.Getenv("ACCOUNTS_SERVICE")), nil)
	if err != nil {
		return utils.ServerError(c, err)
	}

	req.Header.Add("X-Token", c.Get("X-Token"))

	resp, err := conn.Do(req)
	if err != nil {
		return utils.ServerError(c, err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return utils.ServerError(c, err)
	}

	var body map[string]string
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return utils.ServerError(c, err)
	}

	_, exist := body["userID"]
	if !exist {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-token"})
	}

	c.Locals("uuid", body["userID"])
	return c.Next()
}

// func IsLogin(c *fiber.Ctx) error {
// 	claims := getJWTClaims(c)

// 	if claims == nil || claims["uuid"] == nil {
// 		return c.JSON(fiber.Map{"message": "invalid-token"})
// 	}

// 	uuid := claims["uuid"].(string)
// 	c.Locals("uuid", uuid)

// 	return c.Next()
// }

// func getJWTClaims(c *fiber.Ctx) jwt.MapClaims {
// 	jwtToken := c.Get("X-Token")
// 	secret := []byte(os.Getenv("JWT_SECRET"))

// 	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("there was an error")
// 		}
// 		return secret, nil
// 	})

// 	if err != nil || !token.Valid {
// 		return nil
// 	}

// 	return token.Claims.(jwt.MapClaims)
// }
