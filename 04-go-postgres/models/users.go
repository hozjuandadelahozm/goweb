package models

type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

const UserSchema string = `CREATE TABLE users (
	id	SMALLINT NOT NULL PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`
