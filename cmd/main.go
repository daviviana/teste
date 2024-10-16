package main

import (
	"log"
	"teste/internal/adapters/http"
	"teste/internal/config"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	e := http.NewWebService()
	e.Logger.Fatal(e.Start(":8080"))
}
