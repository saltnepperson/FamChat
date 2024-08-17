package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db	*sql.DB


func Initialize(driver, source string) (*sql.DB, error) {
	var err error
	db, err = sql.Open(driver, source)

	if err != nil {
		return nil, fmt.Errorf("Error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Error connecting to database: %v", err)
	}

	log.Println("Connected to database")

	return db, nil
}

func RunDBMigration() {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Errorf("Error creating postgres driver: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	if err != nil {
	fmt.Errorf("Error retrieving migration files: %v", err)
	}
	m.Up()
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
