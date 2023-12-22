package main

import (
	_ "collect-metrics/docs"
	"collect-metrics/handler"
	"collect-metrics/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRouter(metricHandler *handler.MetricHandler) *gin.Engine {
	r := gin.New()
	r.GET("/swag/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/tcp", metricHandler.TCP)
	r.GET("process", metricHandler.Process)
	r.GET("sessions", metricHandler.Session)
	return r
}

func main() {
	newMetricService := service.NewMetricImpl()
	r := SetRouter(handler.NewMetricHandler(newMetricService, "name", 21))
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
