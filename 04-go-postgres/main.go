package main

import (
	"gopostgres/db"
)

func main() {
	db.Connect()
	db.Ping()
	db.Close()
}
