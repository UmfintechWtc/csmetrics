package main

import (
	_ "collect-metrics/docs"
	"collect-metrics/handler"
	"fmt"
	"log"
	"net/http"

	p "collect-metrics/client/prometheus"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetRouter(prom *handler.PrometheusHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.WrapH(promhttp.Handler()))
	r.GET("/tcp", prom.Gauge)
	return r
}

func main() {
	newPrometheus := p.NewMetricsImpl()
	r := SetRouter(handler.NewPrometheusHandler(newPrometheus))
	svr := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: r,
	}
	err := svr.ListenAndServe()
	if err != nil {
		log.Panic(
			fmt.Sprintf("Starting server failed with %s", err),
		)
	}
}
