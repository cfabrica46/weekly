package main

import (
	"log"
	"os"

	"github.com/cfabrica46/weekly/server/internal/all"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("IS_DOCKER") != "true" {
		if err := godotenv.Load(".env"); err != nil {
			log.Println(err)
		}
	}

	r := gin.Default()
	s := r.Group("/api/v1")
	{
		s.GET("/all", all.All)
	}

	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}
