package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
)

var (
	requestProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "inno_taxi_user_number_requests",
		Help: "The total number of processed requests",
	})

	requestSignIn = promauto.NewCounter(prometheus.CounterOpts{
		Name: "inno_taxi_user_number_sign_in_requests",
		Help: "The total number of processed sign-in requests",
	})

	requestSignUp = promauto.NewCounter(prometheus.CounterOpts{
		Name: "inno_taxi_user_number_sign_up_requests",
		Help: "The total number of processed sign-up requests",
	})

	requestSignOut = promauto.NewCounter(prometheus.CounterOpts{
		Name: "inno_taxi_user_number_sign_out_requests",
		Help: "The total number of processed sign-out requests",
	})

	requestGetRating = promauto.NewCounter(prometheus.CounterOpts{
		Name: "inno_taxi_user_number_get_rating_requests",
		Help: "The total number of processed requests of getting rating",
	})
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "inno_taxi_http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})

	httpStatusCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "inno_taxi_http_statuses",
		Help: "The total number of statuses",
	}, []string{"status"})
)

func metricHttpStatus(context *gin.Context) {
	context.Next()
	httpStatusCounter.WithLabelValues(strconv.Itoa(context.Writer.Status())).Inc()
}
