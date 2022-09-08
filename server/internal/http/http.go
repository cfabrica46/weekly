package http

import (
	"log"
	"os"

	"server/internal/all"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	s := r.Group("/api/v1")
	{
		s.GET("/all", all.All)
	}

	log.Println(r.Run(":" + os.Getenv("PORT")))
}
