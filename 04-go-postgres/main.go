package main

import (
	"fmt"
	"gopostgres/db"
)

func main() {
	db.Connect()

	fmt.Println(db.ExistsTable("users"))
	// db.CreateTable(models.UserSchema, "users")

	db.TruncateTable("users")
	db.Close()
	// db.Ping()
}
