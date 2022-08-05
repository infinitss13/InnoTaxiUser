package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/InnoTaxiUser"
	"github.com/infinitss13/InnoTaxiUser/dataBase"
	"github.com/infinitss13/InnoTaxiUser/services"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(context *gin.Context) {
	input := new(InnoTaxiUser.User)
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
	//context.JSON(http.StatusOK, input)
	id, err := dataBase.InsertUser(input)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, "error creating user")
		logrus.Error("status code: ", http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, id)
	logrus.Info("status code: ", http.StatusOK, " User is created")

}

func (h *Handler) signIn(context *gin.Context) {
	input := struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	if err := context.BindJSON(&input); err != nil {

		context.AbortWithStatusJSON(http.StatusBadRequest, "error input data")
		logrus.Error("status code: ", http.StatusBadRequest, " error input data")
		return
	}
	err := dataBase.CheckUser(input.Phone, input.Password)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, "this user doesn't exist")
		logrus.Error("status code: ", http.StatusBadRequest, err)
		return
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
