package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/saltnepperson/FamChat/cmd/server/handler"
	"github.com/saltnepperson/FamChat/util"
)


func main(){
	config, err := util.LoadConfig("../../")

	if err != nil {
		log.Fatalf("Could not load config %v", err)
	} else {
		log.Println("Config file loaded with DB_DRIVER: %v", config.DBDriver)
	}

	fmt.Println("Starting up the web server")
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: handler.RouteService(),
	}

	err = server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
