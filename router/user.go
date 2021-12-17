package router

import (
	"GoFiber/controller"
	"GoFiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(r fiber.Router) {
	router := r.Group("/user")
	router.Post("/register", controller.RegisterUser)
	router.Post("login", controller.LogIn)
	router.Post("/logout", middleware.VerifyUser, controller.LogOut)
}
