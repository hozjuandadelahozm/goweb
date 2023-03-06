package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Guarda la conexión
var db *sql.DB

// Realiza la conexión
func Connect() {
	connStr := "user=root dbname=goweb_db sslmode=verify-full"
	conection, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	fmt.Println("Conexión exitosa")

	db = conection
}

// Cerrar la conexión
func Close() {
	db.Close()
}

// Verificar la conexión
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
