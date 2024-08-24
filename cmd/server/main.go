package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/saltnepperson/FamChat/cmd/server/database"
	"github.com/saltnepperson/FamChat/cmd/server/handler"
	"github.com/saltnepperson/FamChat/cmd/server/middleware"
	"github.com/saltnepperson/FamChat/util"
)

func main() {
	// load configuration
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatalf("Could not load config %v", err)
	}

	// initialize database
	if err = initializeDatabase(config); err != nil {
		log.Fatalf("Failed the initialize database: %w", err)
	}

	// Setup HTTP server
	server := setupServer()

	if err := runServer(server); err != nil {
		log.Fatalf("Server error: %w", err)
	}
}

// loadConfig loads the application configuration
func loadConfig() (util.Config, error) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		return util.Config{}, fmt.Errorf("Could not load config: %w", err)
	}
	log.Println("Config file loaded with DB_DRIVER:", config.DBDriver)
	return config, nil
}

// initialize the database connection and run migrations
func initializeDatabase(config util.Config) error {
	databaseContext, databaseCancel := context.WithCancel(context.Background())
	defer databaseCancel()

	dbConfig := database.Config{
		Driver:          config.DBDriver,
		Source:          config.DBSource,
		MaxOpenConns:    config.DBMaxOpenConns,
		MaxIdleConns:    config.DBMaxIdleConns,
		ConnMaxLifetime: time.Duration(config.DBConnMaxLifetime) * time.Second,
		MigrationPath:   config.DBMigrationPath,
	}

	db, err := database.Initialize(databaseContext, dbConfig)
	if err != nil {
		return fmt.Errorf("Could not connect to the database: %w", err)
	}

	if err := db.RunDBMigration(databaseContext); err != nil {
		return fmt.Errorf("Failed to run database migrations: %w", err)
	}

	return nil
}

// Creates and configures the HTTP server
func setupServer() *http.Server {
	mux := handler.RouteService()
	handler := middleware.Logger(mux)

	return &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}
}

// start the server and handle graceful shutdowns
func runServer(server *http.Server) error {
	serverContext, serverCancel := context.WithCancel(context.Background())
	defer serverCancel()

	// start server go routine
	go func() {
		log.Printf("Starting server on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %w", err)
			serverCancel()
		}
	}()

	// wait for the interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	log.Printf("Shutdown signal recieved. Server shutting down...")

	// create a timetout context for shutdown
	ctx, cancel := context.WithTimeout(serverContext, 30*time.Second)
	defer cancel()

	// attempt a graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("Server forced to shutdown: %w", err)
	}

	log.Printf("Server has shutdown gracefully.")
	return nil
}
