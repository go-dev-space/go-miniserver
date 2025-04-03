package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-dev-space/go-miniserver/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		log.Panic(err)
	}

	system := &system{
		Config: &config{
			Addr: fmt.Sprintf(":%s", os.Getenv("PORT")),
		},
		Handlers: handlers.NewRouteHandlers(),
	}

	log.Printf("[GO SERVER] running => http://localhost%s", system.Config.Addr)
	if err := system.Run(system.Mount()); err != nil {
		log.Panic(err)
	}

}
