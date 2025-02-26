package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type TokenPayload struct {
	Id     string
	Scopes []string
}

func AuthMiddleware(c *fiber.Ctx) error {
	JWTSecretKey := []byte("JWT_SECRET")
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	
	headerParts := strings.Split(tokenString, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		return nil
	}

	newToken := headerParts[1]

	token, err := jwt.Parse(newToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return JWTSecretKey, nil
	})
	if err != nil || !token.Valid {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Token"})
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Email"})

	}

	accountID, _ := claims["Id"].(string)

	if !ok || accountID == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Account"})
	}

	c.Locals("accountId", accountID)

	return c.Next()
}
