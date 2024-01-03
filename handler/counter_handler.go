package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var urlStatusCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "tianciwang",
		Help: "Total number of HTTP requests",
	},
	[]string{"url", "status_code"},
)

func init() {
	// 注册 Counter 到 Prometheus
	prometheus.MustRegister(urlStatusCounter)
}
func (p *PrometheusHandler) Counter(c *gin.Context) {
	fmt.Println("eeeeeeeeeeeeeeeeeeeeeeeeeeee")
	url := c.Request.URL.Path
	statusCode := c.Writer.Status()
	urlStatusCounter.WithLabelValues(url, fmt.Sprintf("%d", statusCode)).Inc()
}
