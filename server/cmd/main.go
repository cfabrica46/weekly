package main

import (
	"log"

	"github.com/cfabrica46/weekly/server/internal/all"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.Use(static.Serve("/", static.LocalFile("./files", false)))
	s := r.Group("/api/v1")
	{
		s.GET("/all", all.All)
	}

	// log.Fatal(r.Run(":" + os.Getenv("PORT")))
	log.Fatal(r.Run(":" + "8080"))
}
