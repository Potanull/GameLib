package main

import (
	"flag"
	"fmt"
	"log"

	"gamelib/internal/server"
	"gamelib/internal/storage/db"
	"gamelib/pkg/config"
	"gamelib/pkg/web"
)

func main() {
	env := flag.String("env", "dev", "env for start")
	flag.Parse()

	if err := config.ParseConfig(*env); err != nil {
		log.Fatalf("[Main] Error initialize configs: %s", err.Error())
	}

	cfg := &config.Config{
		Server:  web.InitServerConfig(),
		Storage: db.InitStorageConfig(),
	}

	srv, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("[Main] Can't connect to database: %s", err.Error())
	} else {
		log.Println("Database connection!")
	}

	path := cfg.Server.Port
	if err := srv.Start(path); err != nil {
		log.Fatalf("[Main] Server start error: %s", err.Error())
	}

	fmt.Print("Hello!")
}
