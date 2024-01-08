package main

import (
	"net/http"
	"os/exec"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	cmd      string = "date"
	result   string
	resultMu sync.Mutex
)

func execCommand() {
	res, err := exec.Command("bash", "-c", cmd).Output()
	if err == nil {
		resultMu.Lock()
		result = string(res)
		resultMu.Unlock()
	}
}

func init() {
	// 在程序运行前先执行一次
	execCommand()
	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-ticker.C:
				res, err := exec.Command("bash", "-c", cmd).Output()
				if err == nil {
					result = string(res)
				}
			}
		}
	}()
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": result})
}
