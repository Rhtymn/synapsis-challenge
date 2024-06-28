package main

import (
	"fmt"
	"net/http"

	"github.com/Rhtymn/synapsis-challenge/config"
	"github.com/Rhtymn/synapsis-challenge/database"
	"github.com/Rhtymn/synapsis-challenge/middleware"
	"github.com/Rhtymn/synapsis-challenge/server"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		fmt.Printf("Init config error: %s\n", err)
		return
	}
	
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Load .env error: %s\n", err)
		return
	}

	db, err := database.ConnectPostgresDB(conf.DatabaseURL)
	if err != nil {
		fmt.Printf("DB Connection Error: %s\n", err)
	}
	defer db.Close()

	corsHandler := middleware.CorsHandler(conf.CorsDomain)

	router := server.SetupServer(server.ServerOpts{
		CorsHandler: corsHandler,
	})
	srv := &http.Server{
		Addr: conf.ServerAddr,
		Handler: router,
	}

	fmt.Printf("Starting Server...\n")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error running server: %v\n", err)
	}
}