package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/innotaxiuser/database"
	"github.com/sirupsen/logrus"
)

func HandleError(err error, context *gin.Context) {
	switch err {
	case database.UserNotFound:
		context.JSON(http.StatusBadRequest, err.Error())
		logrus.Error("status code: ", http.StatusBadRequest, err)
		return
	case database.UserExistErr:
		logrus.Error("status code: ", http.StatusBadRequest, err)
		context.JSON(http.StatusBadRequest, err.Error())
		return

	default:
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		logrus.Error("status code: ", http.StatusInternalServerError, " ", err)
	}

}
func ErrorBinding(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusBadRequest, "wrong inpadsfmiashdfut data")
	logrus.Error("status code: ", http.StatusBadRequest, " wrong inaskdfjjshadfbput data")
}
