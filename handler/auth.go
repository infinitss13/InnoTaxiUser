package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infinitss13/InnoTaxiUser/entity"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (handler AuthHandlers) signUp(context *gin.Context) {

	input := new(entity.User)
	err := context.BindJSON(&input)
	if err != nil {
		handler.loggerMongo.AddNewErrorLog(context, input.Phone, err, "some problems")
		HandleError(err, context)
		return
	}

	id, err := handler.service.CreateUser(*input)
	if err != nil {
		errorCreate := fmt.Errorf("sign-up error, %v", err)
		handler.loggerMongo.AddNewErrorLog(context, input.Phone, errorCreate, "some problems")
		HandleError(err, context)
		return
	}
	err = handler.loggerMongo.AddNewInfoLog(context, input.Phone, "")
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, id)
	logrus.Info("status code: ", http.StatusOK, " User is created")
}

func (handler AuthHandlers) signIn(context *gin.Context) {
	input := new(entity.InputSignIn)
	if err := context.BindJSON(&input); err != nil {
		errorCreate := fmt.Errorf("sign-in error,%v", err)
		handler.loggerMongo.AddNewErrorLog(context, input.Phone, errorCreate, "some problems")
		ErrorBinding(context)
		return
	}
	token, err := handler.service.SignInUser(*input)
	if err != nil {
		errorSignIn := fmt.Errorf("sign-in error : %v", err)
		handler.loggerMongo.AddNewErrorLog(context, input.Phone, errorSignIn, "some problems")
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello",
		"token":   token,
		"time":    time.Now(),
	})
	handler.loggerMongo.AddNewInfoLog(context, input.Phone, "")
	logrus.Info("status code :", http.StatusOK, " user is authorized")
	return
}
