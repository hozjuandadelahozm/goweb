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

	user := models.CreateUser("david", "5678", "david@gmail.com")
	fmt.Println(user)

	// users := models.ListUsers()
	// fmt.Println(users)

	// user := models.GetUser(4)
	// fmt.Println(user)

	// user.Delete()

	// db.TruncateTable("users")
	fmt.Println(models.ListUsers())

	db.Close()
	// db.Ping()
}
