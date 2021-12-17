package middleware

import (
	"GoFiber/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"os"
)

func VerifyUser(c *fiber.Ctx) error {
	secret := os.Getenv("JWT_SECRET")
	tokenString := c.Cookies("token")
	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "You are not login",
		})
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if user, err := model.GetUser(int(claims["user_id"].(float64))); err == nil {
			c.Locals("user", user)
			return c.Next()
		}
		return fiber.ErrNotFound
	} else {
		fmt.Println(err)
		return fiber.ErrUnauthorized
	}
}
