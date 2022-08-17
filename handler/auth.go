package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infinitss13/innotaxiuser/entity"
	"github.com/infinitss13/innotaxiuser/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (handler AuthHandlers) signUp(context *gin.Context) {

	input := new(entity.User)
	err := context.BindJSON(&input)
	if err != nil {
		handler.loggerMongo.LogError(context, err, "some problems")
		HandleError(err, context)
		return
	}

	err = handler.service.CreateUser(*input)
	if err != nil {
		errorCreate := fmt.Errorf("sign-up error, %v", err)
		handler.loggerMongo.LogError(context, errorCreate, "some problems")
		HandleError(err, context)
		return
	}
	err = handler.loggerMongo.LogInfo(context, "")
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, "user successfully created")
	logrus.Info("status code: ", http.StatusOK, " User is created")
}

func (handler AuthHandlers) signIn(context *gin.Context) {
	input := new(entity.InputSignIn)
	if err := context.BindJSON(&input); err != nil {
		errorCreate := fmt.Errorf("sign-in error,%v", err)
		handler.loggerMongo.LogError(context, errorCreate, "some problems")
		ErrorBinding(context)
		return
	}
	token, err := handler.service.SignInUser(*input)
	if err != nil {
		errorSignIn := fmt.Errorf("sign-in error : %v", err)
		handler.loggerMongo.LogError(context, errorSignIn, "some problems")
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, token)
	handler.loggerMongo.LogInfo(context, "")

	logrus.Info("status code :", http.StatusOK, " user is authorized")
	return
}

func (handler AuthHandlers) signOut(context *gin.Context) {

	token, err := middleware.GetToken(context)
	if err != nil {
		errorSignOut := fmt.Errorf("sign-out error: %v", err)
		handler.loggerMongo.LogError(context, errorSignOut, "some problems")
		HandleError(err, context)
		return
	}
	handler.cash.SetValue(token, "true")
	context.JSON(http.StatusOK, "user successfully signed-out")
}
