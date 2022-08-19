package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infinitss13/innotaxiuser/entity"
	"github.com/infinitss13/innotaxiuser/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "inno_taxi_user_number_requests",
		Help: "The total number of processed requests",
	})
	request = promauto.NewCounter(prometheus.CounterOpts{
		Name: "inno_taxi_request_time",
		Help: " request duration",
	})
	buckets               = []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}
	responseTimeHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "namespace",
		Name:      "inno_taxi_user_request_duration",
		Help:      "Histogram of response time for handler in seconds",
		Buckets:   buckets,
	}, []string{"route", "method", "status_code"})
)

func (handler AuthHandlers) signUp(context *gin.Context) {
	startTime := time.Now()
	opsProcessed.Inc()
	input := new(entity.User)
	err := context.BindJSON(&input)
	if err != nil {
		handler.loggerMongo.LogError(context, err)
		ErrorBinding(context)
		return
	}
	err = handler.service.CreateUser(*input)
	if err != nil {
		errorCreate := fmt.Errorf("sign-up error, %v", err)
		handler.loggerMongo.LogError(context, errorCreate)
		HandleError(err, context)
		return
	}
	err = handler.loggerMongo.LogInfo(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, "user successfully created")
	logrus.Info("status code: ", http.StatusOK, " User is created")
	requestTime := time.Since(startTime)
	responseTimeHistogram.WithLabelValues("-", context.Request.Method, "200:OK").Observe(requestTime.Seconds())
	request.Add(requestTime.Seconds())
}

func (handler AuthHandlers) signIn(context *gin.Context) {
	opsProcessed.Inc()
	input := new(entity.InputSignIn)
	if err := context.BindJSON(&input); err != nil {
		errorCreate := fmt.Errorf("sign-in error,%v", err)
		handler.loggerMongo.LogError(context, errorCreate)
		ErrorBinding(context)
		return
	}
	token, err := handler.service.SignInUser(*input)
	if err != nil {
		errorSignIn := fmt.Errorf("sign-in error : %v", err)
		handler.loggerMongo.LogError(context, errorSignIn)
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, token)
	handler.loggerMongo.LogInfo(context)

	logrus.Info("status code :", http.StatusOK, " user is authorized")
	return
}

func (handler AuthHandlers) signOut(context *gin.Context) {
	opsProcessed.Inc()
	token, err := middleware.GetToken(context)
	if err != nil {
		errorSignOut := fmt.Errorf("sign-out error: %v", err)
		handler.loggerMongo.LogError(context, errorSignOut)
		HandleError(err, context)
		return
	}
	handler.cache.SetValue(token, "true")
	context.JSON(http.StatusOK, "user successfully signed-out")
}
