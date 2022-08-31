package all

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func All(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprint(time.Now().Weekday()),
	})
}
