package middleware

import (
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/innotaxiuser/services"
)

func GetToken(context *gin.Context) (string, error) {
	tokenString := context.GetHeader("Authorization")
	if tokenString == "" {
		context.JSON(401, gin.H{"error": "request does not contain an access token"})
		context.Abort()
		return "", errors.New("no access token")
	}
	splitedToken := strings.Split(tokenString, " ")
	return splitedToken[1], nil
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		splitedToken, err := GetToken(context)
		if err != nil {
			logrus.Error(err)
			context.JSON(http.StatusInternalServerError, err)
		}
		_, err = services.VerifyToken(splitedToken)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
	}
}
