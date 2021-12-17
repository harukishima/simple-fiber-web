package controller

import (
	"GoFiber/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func RegisterUser(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	fmt.Println(name, email, password)
	if email == "" || len([]rune(password)) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
		})
	}
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server error",
		})
	} else {
		user := model.User{Name: name, Email: email, Password: string(hashedPassword)}
		if err2 := model.CreateUser(user); err2 != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Server error2",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User created",
		})
	}
}

func LogIn(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	secret := os.Getenv("JWT_SECRET")
	user, err := model.GetUserByEmail(email)
	if err != nil {
		return fiber.ErrNotFound
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Password do not match",
		})
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.Id,
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	cookie := &fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(cookie)
	return c.JSON(fiber.Map{
		"message": "Log in success",
	})
}

func LogOut(c *fiber.Ctx) error {
	cookie := &fiber.Cookie{
		Name:    "token",
		Expires: time.Now().Add(-24 * time.Hour),
	}
	c.Cookie(cookie)
	return c.Status(200).JSON(fiber.Map{
		"message": "Logout success",
	})
}
