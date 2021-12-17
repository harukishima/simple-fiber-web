package model

import (
	"GoFiber/db"
	"fmt"
)

type Restaurant struct {
	Id      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Point   int    `json:"point,omitempty"`
}

func AddNewRestaurant(restaurant Restaurant) error {
	_, err := db.DB.NamedExec("insert into restaurant(name, address, point) values (:name, :address, :point)", restaurant)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetAllRestaurant() []Restaurant {
	restaurants := []Restaurant{}
	err := db.DB.Select(&restaurants, "select * from restaurant")
	if err != nil {
		return nil
	}
	return restaurants
}

func GetRestaurant(id int) (Restaurant, error) {
	restaurant := Restaurant{}
	err := db.DB.Get(&restaurant, "select * from restaurant where ?", id)
	if err != nil {
		return Restaurant{}, err
	}
	return restaurant, nil
}

func DeleteRestaurant(id int) error {
	exec, err := db.DB.Exec("delete from restaurant where id = $1", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rows, _ := exec.RowsAffected()
	fmt.Println(rows, "was removed")
	return nil
}
