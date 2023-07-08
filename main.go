package main

import (
	"github.com/betelgeusexru/golang-prometheus-grafana/internal/middlewares"
	"github.com/betelgeusexru/golang-prometheus-grafana/pkg/metrics"
	"github.com/gin-gonic/gin"
)

func main() {
	metrics.InitCustomMetrics()

	r := gin.Default()

	r.Use(middlewares.IncRequests())

	r.GET("/hello", func(c *gin.Context) {
		metrics.GetMetrics().HelloEndpointHitCount.Inc()
		c.String(200, "hello")
	})

	r.GET("/metrics", metrics.PrometheusHandler())

	r.Run(":3000")
}
