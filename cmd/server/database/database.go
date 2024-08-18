package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var db *sql.DB

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
		log.Fatalf("Error creating postgres driver: %v", err)
	}
	log.Printf("Driver value: %+v", driver)

	m, err := migrate.NewWithDatabaseInstance(
		"file:///famchat/db/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Error initializing migration: %v", err)
	}
	log.Printf("Migration instance: %+v", m)

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not apply migration: %v", err)
	}

	log.Println("Successfully applied migrations...")
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
