package handler

import "github.com/gin-gonic/gin"

// 实例化TcpMetric
func (i *MetricHandler) TCP(c *gin.Context) {
	err := i.service.TCP("tcp", 21)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "err tcp",
		})
	}
	c.JSON(200, gin.H{
		"message": "tcp",
	})
	return
}

// 实例化SessionMetric
func (i *MetricHandler) Session(c *gin.Context) {
	err := i.service.TCP("session", 21)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "err session",
		})

	}
	c.JSON(200, gin.H{
		"message": "session",
	})
	return
}

// 实例化ProcessMetric
func (i *MetricHandler) Process(c *gin.Context) {
	err := i.service.TCP("process", 21)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "err process",
		})

	}
	c.JSON(200, gin.H{
		"message": "process",
	})
	return
}
