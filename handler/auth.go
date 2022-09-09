package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/innotaxiuser/entity"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

// @Summary SignUp
// @Tags auth
// @Description handler for SignUp request, allows user to register in service
// @ID signup
// @Param input body entity.User true "account info"
// @Accept json
// @Produce json
// @Success 200 {string} string "user successfully created"
// @Failure 400 {object} error
// @Router /auth/sign-up [post]
func (handler AuthHandlers) signUp(context *gin.Context) {
	timer := prometheus.NewTimer(httpDuration.WithLabelValues(context.Request.RequestURI))
	requestProcessed.Inc()
	requestSignUp.Inc()
	input := new(entity.User)
	err := context.BindJSON(&input)
	if err != nil {
		if errorLogger := handler.LoggerMongo.LogError(context, err); errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger.Error())
		}
		ErrorBinding(context)
		return
	}
	err = handler.UserService.CreateUser(*input)

	if err != nil {
		errorCreate := fmt.Errorf("sign-up error, %v", err)
		if errorLogger := handler.LoggerMongo.LogError(context, errorCreate); errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger.Error())
		}
		HandleError(err, context)
		return
	}
	err = handler.LoggerMongo.LogInfo(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, "user successfully created")
	logrus.Info("status code: ", http.StatusOK, " User is created")
	timer.ObserveDuration()

}

// @Summary SignIn
// @Tags auth
// @Description handler for SignIn request, allows user to authenticate
// @Param input body entity.InputSignIn true "account info"
// @Accept json
// @Produce json
// @Success 200 {string} token "token"
// @Failure 400 {object} error
// @Router /auth/sign-in [post]
func (handler AuthHandlers) signIn(context *gin.Context) {
	timer := prometheus.NewTimer(httpDuration.WithLabelValues(context.Request.RequestURI))
	requestProcessed.Inc()
	requestSignIn.Inc()
	input := new(entity.InputSignIn)
	if err := context.BindJSON(&input); err != nil {
		errorCreate := fmt.Errorf("sign-in error,%v", err)
		errorLogger := handler.LoggerMongo.LogError(context, errorCreate)
		if errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger.Error())
		}
		ErrorBinding(context)
		return
	}
	token, err := handler.UserService.SignInUser(*input)
	if err != nil {
		errorSignIn := fmt.Errorf("sign-in error : %v", err)
		if errorLogger := handler.LoggerMongo.LogError(context, errorSignIn); errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger.Error())
		}
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, token)
	if err = handler.LoggerMongo.LogInfo(context); err != nil {
		logrus.Error("error in logger : ", err)
	}
	logrus.Info("status code :", http.StatusOK, " user is authorized")
	timer.ObserveDuration()

}

// @Summary SignOut
// @Security ApiKeyAuth
// @Tags auth
// @Description handler for SignOut request, allows user log out of his account
// @Produce json
// @Success 200 {string} message "user successfully signed-out"
// @Failure 400 {object} error
// @Router /auth/sign-out [get]
func (handler AuthHandlers) signOut(context *gin.Context) {
	timer := prometheus.NewTimer(httpDuration.WithLabelValues(context.Request.RequestURI))
	requestProcessed.Inc()
	requestSignOut.Inc()
	token, err := handler.UserService.GetToken(context)
	if err != nil {
		errorSignOut := fmt.Errorf("sign-out error: %v", err)
		if errorLogger := handler.LoggerMongo.LogError(context, errorSignOut); errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger.Error())
		}
		HandleError(err, context)
		return
	}
	if err = handler.Cache.SetValue(token, "true"); err != nil {
		errorSignOut := fmt.Errorf("sign-out error: %v", err)
		if errorLogger := handler.LoggerMongo.LogError(context, errorSignOut); errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger.Error())
		}
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, "user successfully signed-out")
	timer.ObserveDuration()

}
