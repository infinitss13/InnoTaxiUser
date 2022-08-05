package handler

import (
	"database/sql"
	"errors"
	"net/http"

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
		logrus.Error("status code: ", http.StatusBadRequest, err)
		context.AbortWithStatusJSON(http.StatusBadRequest, "error input data")
		return
	}

	input, err = services.CreateUser(input)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, "error creating user")
		logrus.Error("status code: ", http.StatusBadRequest, err)
		return

	}
	var inputError = errors.New("user exists")
	id, err := dataBase.InsertUser(configs.NewConfig(), input)
	if err != nil {
		if err.Error() == inputError.Error() {
			logrus.Error("status code: ", http.StatusBadRequest, err)
			context.AbortWithStatusJSON(http.StatusBadRequest, "user with this data already exists")
			return
		} else {
			context.AbortWithStatusJSON(http.StatusInternalServerError, "error creating user")
			logrus.Error("status code: ", http.StatusInternalServerError, err)
			return
		}
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
		context.AbortWithStatusJSON(http.StatusBadRequest, "error input data")
		logrus.Error("status code: ", http.StatusBadRequest, " error input data")
		return
	}

	err := dataBase.CheckUser(configs.NewConfig(), input.Phone, input.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.AbortWithStatusJSON(http.StatusBadRequest, "user doesn't exist")
			logrus.Error("status code: ", http.StatusBadRequest, "user doesn't exist")
			return
		} else {
			context.AbortWithStatusJSON(http.StatusInternalServerError, err)
			logrus.Error("status code: ", http.StatusInternalServerError, err)
			return
		}
	}

	token, err := services.CreateToken(input.Phone)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		logrus.Error("status code: ", http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello",
		"token":   token,
	})
	logrus.Info("status code :", http.StatusOK, " user is authorized")

}
