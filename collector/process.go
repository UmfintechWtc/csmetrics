package collector

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ProcessGuage(c *gin.Context) {
	fmt.Println("TianCiwang")
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
