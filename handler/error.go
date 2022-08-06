package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/infinitss13/InnoTaxiUser/dataBase"
	"github.com/sirupsen/logrus"
	"net/http"
)

func HandleError(err error, context *gin.Context) {
	switch err {
	case dataBase.UserNotFound:
		context.AbortWithStatusJSON(http.StatusBadRequest, " user doesn't exist")
		logrus.Error("status code: ", http.StatusBadRequest, " user doesn't exist")
		return
	case dataBase.UserExistErr:
		logrus.Error("status code: ", http.StatusBadRequest, err)
		context.AbortWithStatusJSON(http.StatusBadRequest, "user with this data already exists")
		return

	default:
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		logrus.Error("status code: ", http.StatusInternalServerError, " ", err)
	}

}
func ErrorBinding(err error, context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusBadRequest, "wrong input data")
	logrus.Error("status code: ", http.StatusBadRequest, " wrong input data")
	return
}
