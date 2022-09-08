package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("IS_DOCKER") != "true" {
		if err := godotenv.Load(".env"); err != nil {
			log.Println(err)
		}
	}
}
