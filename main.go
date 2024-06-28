package main

import (
	"fmt"

	"github.com/Rhtymn/synapsis-challenge/config"
	"github.com/Rhtymn/synapsis-challenge/database"
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
}