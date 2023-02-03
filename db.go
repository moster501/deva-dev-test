package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type (
	connString struct {
		server, port, databaseName, user, password string
	}
)

var DBMySQL = connect(&connString{
	server:       GetEnvConfig("Server1"),
	port:         GetEnvConfig("Port1"),
	databaseName: GetEnvConfig("Database1"),
	user:         GetEnvConfig("User1"),
	password:     GetEnvConfig("Password1"),
})

func connect(info *connString) *sql.DB {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", info.user, info.password, info.server, info.port, info.databaseName)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return db
}
