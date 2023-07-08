package middlewares

import (
	"strings"

	"github.com/betelgeusexru/golang-prometheus-grafana/pkg/metrics"
	"github.com/gin-gonic/gin"
)

func IncRequests() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !strings.Contains(c.Request.URL.String(), "metrics") {
			metrics.GetMetrics().TotalAmountOfRequest.Inc()
		}

		c.Next()
	}
}
