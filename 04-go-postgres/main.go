package main

import (
	"fmt"
	"gopostgres/db"
	"gopostgres/models"
)

func main() {
	db.Connect()

	// fmt.Println(db.ExistsTable("users"))
	// db.CreateTable(models.UserSchema, "users")

	// user := models.CreateUser("rosmeri", "9635", "rosmeri@gmail.com")
	// fmt.Println(user)

	// users := models.ListUsers()
	// fmt.Println(users)

	user := models.GetUser(4)
	fmt.Println(user)

	// db.TruncateTable("users")
	db.Close()
	// db.Ping()
}
