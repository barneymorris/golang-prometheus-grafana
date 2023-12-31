package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type tCustomMetrics struct {
	HelloEndpointHitCount prometheus.Counter
	TotalAmountOfRequest  prometheus.Counter
}

var customMetrics *tCustomMetrics

func InitCustomMetrics() {
	reg := prometheus.NewRegistry()
	customMetrics = &tCustomMetrics{
		HelloEndpointHitCount: promauto.NewCounter(prometheus.CounterOpts{
			Name: "hello_endpoint_hitted",
			Help: "The total number of GET /hello endpoint hitted",
		}),
		TotalAmountOfRequest: promauto.NewCounter(prometheus.CounterOpts{
			Name: "total_amount_of_request",
			Help: "The total amount of all request",
		}),
	}

	reg.MustRegister(
		customMetrics.HelloEndpointHitCount,
		customMetrics.TotalAmountOfRequest,
	)

}

func GetMetrics() *tCustomMetrics {
	return customMetrics
}

func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
