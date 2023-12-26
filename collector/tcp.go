package collector

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// TianCiwang godoc
// @Summary Your endpoint summary
// @Description Your endpoint description
// @Tags Your endpoint tags
// @Accept json
// @Produce json
// @Router /test [get]
func TcpGuage(c *gin.Context) {
	fmt.Println("TianCiwang")
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
