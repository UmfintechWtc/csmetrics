package debug

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

var cmd string = "date"

func execCommand(cmdRes chan string) error {
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			res, err := exec.Command("bash", "-c", cmd).Output()
			if err != nil {
				return nil
			}
			cmdRes <- string(res)
		}
	}
}

func main() {
	cmdRes := make(chan string)
	go func() {
		execCommand(cmdRes)
	}()
	go func() {
		for {
			select {
			case result := <-cmdRes:
				fmt.Println(fmt.Sprintf("current time: %v", result))
			default:
				fmt.Println("the chan is empty")
			}
			time.Sleep(3 * time.Second)
		}
	}()
	// 设置GIN路由
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// 启动GIN HTTP服务
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting HTTP server:", err)
	}
}
