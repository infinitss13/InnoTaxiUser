package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/InnoTaxiUser/dataBase"
	"github.com/sirupsen/logrus"
)

func HandleError(err error, context *gin.Context) {
	switch err {
	case dataBase.UserNotFound:
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		logrus.Error("status code: ", http.StatusBadRequest, err)
		return
	case dataBase.UserExistErr:
		logrus.Error("status code: ", http.StatusBadRequest, err)
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return

	default:
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		logrus.Error("status code: ", http.StatusInternalServerError, " ", err)
	}

}
func ErrorBinding(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusBadRequest, "wrong input data")
	logrus.Error("status code: ", http.StatusBadRequest, " wrong input data")
	return
}
