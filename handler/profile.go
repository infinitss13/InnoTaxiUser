package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/innotaxiuser/entity"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

func (handler AuthHandlers) getProfile(context *gin.Context) {
	requestGetProfile.Inc()
	requestProcessed.Inc()
	timer := prometheus.NewTimer(httpDuration.WithLabelValues(context.Request.RequestURI))

	tokenSigned, err := handler.getAndCheckToken(context)
	if err != nil {
		errorToken := fmt.Errorf("profile error,%v", err)
		if errorLogger := handler.loggerMongo.LogError(context, errorToken); errorLogger != nil {
			logrus.Error("error in logger", errorLogger)
		}
		HandleError(err, context)
		return
	}
	userData, err := handler.service.GetUserByToken(tokenSigned)
	if err != nil {
		errorToken := fmt.Errorf("profile error,%v", err)
		if errorLogger := handler.loggerMongo.LogError(context, errorToken); errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger.Error())
		}
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, userData)
	if err = handler.loggerMongo.LogInfo(context); err != nil {
		logrus.Error("error in logger : ", err)
	}
	logrus.Info("status code :", http.StatusOK, " user get profile")
	timer.ObserveDuration()

}

func (handler AuthHandlers) updateProfile(context *gin.Context) {
	update := new(entity.UpdateData)
	if err := context.BindJSON(&update); err != nil {
		errorCreate := fmt.Errorf("update profile error: %v", err)
		handler.loggerMongo.LogError(context, errorCreate)
		ErrorBinding(context)
		return
	}

	tokenSigned, err := handler.getAndCheckToken(context)
	if err != nil {
		errorToken := fmt.Errorf("update profile error,%v", err)
		handler.loggerMongo.LogError(context, errorToken)
		HandleError(err, context)
		return
	}

	err = handler.service.UpdateUserProfile(tokenSigned, update)
	if err != nil {
		errorUpdate := fmt.Errorf("update profile error,%v", err)
		handler.loggerMongo.LogError(context, errorUpdate)
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, "User is successfully updated")
	logrus.Info("status code :", http.StatusOK, " user updated profile")
	if err = handler.cache.SetValue(tokenSigned, "true"); err != nil {
		logrus.Error("error in cache db : ", err.Error())
		handler.loggerMongo.LogError(context, err)
	}
	if err = handler.loggerMongo.LogInfo(context); err != nil {
		logrus.Error("error in logger : ", err)
	}
}

func (handler AuthHandlers) deleteProfile(context *gin.Context) {
	deleteData := new(entity.DeleteData)
	if err := context.BindJSON(&deleteData); err != nil {
		errorCreate := fmt.Errorf("delete profile error: %v", err)
		handler.loggerMongo.LogError(context, errorCreate)
		ErrorBinding(context)
		return
	}
	tokenSigned, err := handler.getAndCheckToken(context)
	if err != nil {
		errorToken := fmt.Errorf("delete profile error: %v", err)
		handler.loggerMongo.LogError(context, errorToken)
		HandleError(err, context)
		return
	}
	err = handler.service.DeleteProfile(tokenSigned, deleteData.Password)
	if err != nil {
		errorToken := fmt.Errorf("delete profile error: %v", err)
		errorLog := handler.loggerMongo.LogError(context, errorToken)
		if errorLog != nil {
			logrus.Info("error in logger : ", errorLog.Error())
		}
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, "User is successfully deleted")
	logrus.Info("status code :", http.StatusOK, " user deleted profile")
	if err = handler.cache.SetValue(tokenSigned, "true"); err != nil {
		logrus.Error("error in cache db : ", err.Error())
		handler.loggerMongo.LogError(context, err)
	}
	if err = handler.loggerMongo.LogInfo(context); err != nil {
		logrus.Error("error in logger : ", err.Error())
	}
}
