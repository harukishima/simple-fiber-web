package model

import (
	"GoFiber/db"
	"fmt"
)

type User struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email"`
	Type     int    `json:"type,omitempty"`
	Password string `json:"-"`
}

func CreateUser(user User) error {
	if _, err := db.DB.NamedExec("insert into users(name, email, password) values (:name, :email, :password)", user); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetUser(id int) (User, error) {
	user := User{}
	if err := db.DB.Get(&user, "select * from users where id = $1", id); err == nil {
		fmt.Println(err)
		return User{}, err
	}
	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	user := User{}
	if err := db.DB.Get(&user, "select * from users where email = $1", email); err != nil {
		fmt.Println(err)
		return User{}, err
	}
	return user, nil
}

func DeleteUser(id int) error {
	if _, err := db.DB.Exec("delete from users where id = $1", id); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
