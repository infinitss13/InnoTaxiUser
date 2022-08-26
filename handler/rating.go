package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/innotaxiuser/cmd/cache"
	"github.com/infinitss13/innotaxiuser/middleware"
	"github.com/sirupsen/logrus"
)

type getRate struct {
	Phone  string
	Rating float32
}

func (handler AuthHandlers) getRating(context *gin.Context) {
	requestProcessed.Inc()
	requestGetRating.Inc()
	tokenSigned, err := middleware.GetToken(context)
	if err != nil {
		errorRating := fmt.Errorf("get rating error : %v", err)
		if errorLogger := handler.loggerMongo.LogError(context, errorRating); errorLogger != nil {
			logrus.Info("error in logger : ", errorLogger)
		}
		HandleError(err, context)
		return
	}
	isKey, err := handler.cache.GetValue(tokenSigned)
	if isKey && err != cache.UserSignedOut {
		context.JSON(http.StatusBadRequest, "user with this token signed-out")
		return
	}
	rating, userPhone, err := handler.service.GetRatingWithToken(tokenSigned)
	if err != nil {
		errorRating := fmt.Errorf("get rating error : %v", err)
		if errorLogger := handler.loggerMongo.LogError(context, errorRating); errorLogger != nil {
			logrus.Info("error in logger : ", errorLogger.Error())
		}
		HandleError(err, context)
		return
	}
	rate := getRate{
		Phone:  userPhone,
		Rating: rating,
	}
	context.JSON(http.StatusOK, rate)
	if errorLogger := handler.loggerMongo.LogInfo(context); errorLogger != nil {
		logrus.Error("error in logger : ", errorLogger.Error())
	}
	logrus.Info("status code :", http.StatusOK, " user get rating")

}

type ratingTrip struct {
	Rating int `json:"rating"`
}

func rateTrip(context *gin.Context) {
	inputRating := new(ratingTrip)
	err := context.BindJSON(&inputRating)
	if err != nil {
		ErrorBinding(context)
		return
	}
	// TODO : ere will me method that address to order service`s database and insert rating to the last trip of user, if time is not expired
}
