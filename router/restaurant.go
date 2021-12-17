package router

import (
	"GoFiber/controller"
	"GoFiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func RestaurantRoute(r fiber.Router) {
	router := r.Group("/restaurant")
	router.Get("/", controller.GetAllRestaurant)
	router.Get("/:id", controller.GetRestaurant)
	router.Post("/", middleware.VerifyUser, controller.AddRestaurant)
	router.Delete("/:id", middleware.VerifyUser, controller.DeleteRestaurant)
}
