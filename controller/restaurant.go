package controller

import (
	"GoFiber/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetAllRestaurant(c *fiber.Ctx) error {
	restaurants := model.GetAllRestaurant()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": restaurants,
	})
}

func GetRestaurant(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if restaurant, err := model.GetRestaurant(id); err == nil {
		if id == restaurant.Id {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": restaurant,
			})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Not found",
	})
}

func AddRestaurant(c *fiber.Ctx) error {
	var restaurant = new(model.Restaurant)
	if err := c.BodyParser(restaurant); err == nil {
		if err2 := model.AddNewRestaurant(*restaurant); err2 != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Cannot insert",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Restaurant added",
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "Not supported",
	})
}

func DeleteRestaurant(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := model.DeleteRestaurant(id); err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Delete success",
		})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Not found",
	})
}
