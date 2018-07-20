package db

import (
	"os"
	"fmt"
	"database/sql"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database"
)

var db *sql.DB

func Connect() {
	var driver database.Driver
	var err error

	dbUrl := os.Getenv("DATABASE_URL")
	fmt.Println(dbUrl)
	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		println(1, err.Error())
		return
	}

	driver, err = postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		println(2, err.Error())
		return
	}

	migrateDatabase(driver)
}

func Get() *sql.DB {
	return db
}

func migrateDatabase(driver database.Driver) {
	println("Performing migration")
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)

	if err != nil {
		println("Error on connect:", err.Error())
		return
	}

	err = m.Up()

	if err != nil {
		println("Error on up:", err.Error())
		return
	}
}