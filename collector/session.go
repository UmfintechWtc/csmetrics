package collector

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SessionGuage(c *gin.Context) {
	fmt.Println("TianCiwang")
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
