package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/InnoTaxiUser/configs"
	"github.com/infinitss13/InnoTaxiUser/dataBase"
	"github.com/infinitss13/InnoTaxiUser/entity"
	"github.com/infinitss13/InnoTaxiUser/services"
	"github.com/sirupsen/logrus"
)

func signUp(context *gin.Context) {
	input := new(entity.User)
	err := context.BindJSON(&input)
	if err != nil {
		HandleError(err, context)
		return
	}

	input, err = services.CreateUser(input)
	if err != nil {
		HandleError(err, context)
		return
	}

	db, err := dataBase.NewDB(configs.NewConfig())
	id, err := db.InsertUser(input)

	if err != nil {
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, id)
	logrus.Info("status code: ", http.StatusOK, " User is created")
}

func signIn(context *gin.Context) {
	input := struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := context.BindJSON(&input); err != nil {
		ErrorBinding(err, context)
		return
	}
	db, err := dataBase.NewDB(configs.NewConfig())
	if err != nil {
		HandleError(err, context)
		return
	}
	err = db.UserIsRegistered(input.Phone, input.Password)
	if err != nil {
		HandleError(err, context)
		return
	}
	token, err := services.CreateToken(input.Phone)
	if err != nil {
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello",
		"token":   token,
		"time":    time.Now(),
	})
	logrus.Info("status code :", http.StatusOK, " user is authorized")

}
