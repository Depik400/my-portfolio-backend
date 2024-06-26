package main

import (
	"log"
	"pavelkononov/resume/internal/handlers"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	handlers.MakeHttpServer()
}
