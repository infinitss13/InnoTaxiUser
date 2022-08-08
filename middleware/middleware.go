package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/infinitss13/InnoTaxiUser/services"
	"strings"
)

func GetToken(ctx *gin.Context) (string, error) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.JSON(401, gin.H{"error": "request does not contain an access token"})
		ctx.Abort()
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
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		//context.JSON(http.StatusOK, phone)

	}
}
