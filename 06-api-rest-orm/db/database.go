package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Realizar la conexión
var dsn = "host=localhost user=root password=root dbname=goweb_db port=5505 sslmode=disable TimeZone=America/Bogota"
var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la conxión", err)
		panic(err)
	} else {
		fmt.Println("Conexión exitosa")
		return db
	}
}()
