package main

import (
	"GoFiber/db"
	"GoFiber/router"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
	db.ConnectDb()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover2.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "hello-world",
		})
	})
	router.RestaurantRoute(app)
	router.UserRoute(app)
	app.Get("/err", func(c *fiber.Ctx) error {
		panic("I'm an error")
	})
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
	if err := app.Listen(":8080"); err != nil {
		log.Println("Can not start server")
	}
}
