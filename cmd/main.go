package main

import (
	"database/sql"
	"log"

	"github.com/sila1404/go-http-standard-lib/cmd/api"
	"github.com/sila1404/go-http-standard-lib/config"
	"github.com/sila1404/go-http-standard-lib/db"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMyStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer("localhost:8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// initStorage initializes the storage by establishing a connection to the database.
//
// Parameters:
// - db: a pointer to the sql.DB object representing the database connection.
//
// Return type: None.
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Connection established successfully")
}
