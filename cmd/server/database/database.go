package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// db.sql configuration with additional configurations
type DB struct {
	*sql.DB
	config Config
}

// holds database configuration
type Config struct {
	Driver          string
	Source          string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	MigrationPath   string
}

// creates a new connection and applies configuration
func Initialize(ctx context.Context, config Config) (*DB, error) {
	db, err := sql.Open(config.Driver, config.Source)
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %v", err)
	}

	if err = db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("Error connecting to database: %v", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)

	log.Println("Connected to database")

	return &DB{DB: db, config: config}, nil
}

// applies database migrations
func (db *DB) RunDBMigration(ctx context.Context) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("Error creating postgres driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		db.config.MigrationPath,
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("Error initializing migration: %v", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("Could not apply migration: %v", err)
	}

	return nil
}

// close the database connection
func (db *DB) Close() error {
	if db.DB != nil {
		return db.DB.Close()
	}
	return nil
}
