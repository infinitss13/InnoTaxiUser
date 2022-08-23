package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (handler AuthHandlers) getProfile(context *gin.Context) {
	requestGetProfile.Inc()
	requestProcessed.Inc()
	timer := prometheus.NewTimer(httpDuration.WithLabelValues(context.Request.RequestURI))

	tokenSigned, err := handler.getAndCheckToken(context)
	if err != nil {
		errorToken := fmt.Errorf("profile error,%v", err)
		handler.loggerMongo.LogError(context, errorToken)
		HandleError(err, context)
		return
	}

	userData, err := handler.service.GetUserByToken(tokenSigned)
	if err != nil {
		errorToken := fmt.Errorf("profile error,%v", err)
		handler.loggerMongo.LogError(context, errorToken)
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, userData)
	handler.loggerMongo.LogInfo(context)
	logrus.Info("status code :", http.StatusOK, " user get profile")
	timer.ObserveDuration()

}
