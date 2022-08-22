package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infinitss13/innotaxiuser/entity"
	"github.com/infinitss13/innotaxiuser/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (handler AuthHandlers) signUp(context *gin.Context) {
	timer := prometheus.NewTimer(httpDuration.WithLabelValues(context.Request.RequestURI))
	requestProcessed.Inc()
	requestSignUp.Inc()
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
	//httpStatusCounter.WithLabelValues("http.StatusOK").Inc()
	logrus.Info("status code: ", http.StatusOK, " User is created")
	timer.ObserveDuration()

}

func (handler AuthHandlers) signIn(context *gin.Context) {
	requestProcessed.Inc()
	requestSignIn.Inc()
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
	//httpStatusCounter.WithLabelValues("http.StatusOK").Inc()
	logrus.Info("status code :", http.StatusOK, " user is authorized")
	return
}

func (handler AuthHandlers) signOut(context *gin.Context) {
	requestProcessed.Inc()
	requestSignOut.Inc()
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
