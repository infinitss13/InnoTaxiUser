package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/InnoTaxiUser/services"
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
		_, err = services.VerifyToken(splitedToken)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
	}
}
