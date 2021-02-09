package main

import (
	"log"

	"cmd/src/main.go/lib/services"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	services.StartAllServices()
}
