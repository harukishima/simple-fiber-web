package db

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDb() {
	var err error
	DB, err = sqlx.Connect("pgx", "postgres://postgres:newadminpassword@localhost:5432/gofiber")
	if err != nil {
		fmt.Println("connect DB fail: ", err)
		return
	}
	fmt.Println("connect DB success")
}
