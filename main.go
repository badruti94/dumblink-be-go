package main

import (
	"dumblink-be-go/src/routes"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := routes.Router()

	router.Run(":8080")
}
