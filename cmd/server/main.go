package main

import (
	"fmt"
	"log"
	"net/http"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/saltnepperson/FamChat/cmd/server/handler"
	"github.com/saltnepperson/FamChat/cmd/server/middleware"
	"github.com/saltnepperson/FamChat/util"
)


func main(){
	mux := handler.RouteService()
	config, err := util.LoadConfig("../../")

	if err != nil {
		log.Fatalf("Could not load config %v", err)
	} else {
		log.Println("Config file loaded with DB_DRIVER: %s", config.DBDriver)
	}

	fmt.Println("Starting up the web server...")

	handler := middleware.Logger(mux)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: handler,
	}

	serverContext, serverCancel := context.WithCancel(context.Background())
	defer serverCancel()

	go handleGracefulShutdown(serverContext, server, serverCancel)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Block until the context is Done
	<-serverContext.Done()
	log.Println("Server has shutdown...")
}

func handleGracefulShutdown(cxt context.Context, server *http.Server, cancel context.CancelFunc) {
	// Listen for interrupt/cancel signal
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Block until the signal is received
	<-signalChannel

	// Create a context with timeout for the shutdwon process
	shutdownContext, shutdownCancel := context.WithTimeout(cxt, 30*time.Second)
	defer shutdownCancel()

	log.Println("Shutdown signal received. Server shutting down...")

	if err := server.Shutdown(shutdownContext); err != nil {
		log.Printf("Server shutdown failed: %v", err)
	}

	// If the shutdown exceeds the timeout, log and forcefully shutdown
	<-shutdownContext.Done()
	if shutdownContext.Err() == context.DeadlineExceeded {
		log.Fatal("Graceful shutdown timeout exceeded... forcing an exit.")
	}

	cancel()
}

