package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/InnoTaxiUser/middleware"
	"github.com/sirupsen/logrus"
)

type getRate struct {
	Phone  string
	Rating float32
}

func (handler AuthHandlers) getRating(context *gin.Context) {
	tokenSigned, err := middleware.GetToken(context)
	if err != nil {
		errorRating := fmt.Errorf("get rating error : %v", err)
		handler.loggerMongo.AddNewErrorLog(context, "-", errorRating, "some problems")
		HandleError(err, context)
		return
	}
	rating, userPhone, err := handler.service.GetRatingWithToken(tokenSigned)
	if err != nil {
		errorRating := fmt.Errorf("get rating error : %v", err)
		handler.loggerMongo.AddNewErrorLog(context, userPhone, errorRating, "some problems")
		HandleError(err, context)
		return
	}
	rate := getRate{
		Phone:  userPhone,
		Rating: rating,
	}
	context.JSON(http.StatusOK, rate)
	handler.loggerMongo.AddNewInfoLog(context, userPhone, "")
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
