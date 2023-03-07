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
	// connStr := "user=root password=root host=localhost port=5505 dbname=goweb_db sslmode=verify-full"
	connStr := "user=root password=root host=localhost port=5505 dbname=goweb_db sslmode=disable"
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

// Verifica si una tabla existe o no
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf(`
    select
    table_name
    from information_schema.tables
    where table_name like '%s'
       and table_schema not in ('information_schema', 'pg_catalog')
       and table_type = 'BASE TABLE'
    order by table_name,
    table_schema;`,
		tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return rows.Next()
}

// Crea una tabla
func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}

}
