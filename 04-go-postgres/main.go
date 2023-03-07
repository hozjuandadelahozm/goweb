package main

import (
	"fmt"
	"gopostgres/db"
	"gopostgres/models"
)

func main() {
	db.Connect()

	fmt.Println(db.ExistsTable("users"))
	db.CreateTable(models.UserSchema, "users")

	db.Close()
	// db.Ping()
}
