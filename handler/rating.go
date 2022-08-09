package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/infinitss13/InnoTaxiUser/configs"
	"github.com/infinitss13/InnoTaxiUser/dataBase"
	"github.com/infinitss13/InnoTaxiUser/middleware"
	"github.com/infinitss13/InnoTaxiUser/services"
	"net/http"
)

func getRating(context *gin.Context) {
	tokenSigned, err := middleware.GetToken(context)
	if err != nil {
		HandleError(err, context)
		return
	}
	db, err := dataBase.NewDB(configs.NewConfig())
	if err != nil {
		HandleError(err, context)
		return
	}
	claims, err := services.VerifyToken(tokenSigned)
	if err != nil {
		HandleError(err, context)
		return
	}
	userPhone := claims.Phone
	rating, err := db.GetRatingByPhone(userPhone)
	if err != nil {
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"phone":  userPhone,
		"rating": rating,
	})

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
	//here will me method that address to order service`s database and insert rating to the last trip of user, if time is not expired

}
