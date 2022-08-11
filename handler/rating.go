package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/InnoTaxiUser/configs"
	"github.com/infinitss13/InnoTaxiUser/dataBase"
	"github.com/infinitss13/InnoTaxiUser/dataBase/mongoDB"
	"github.com/infinitss13/InnoTaxiUser/middleware"
	"github.com/infinitss13/InnoTaxiUser/services"
	"github.com/sirupsen/logrus"
)

func getRating(context *gin.Context) {
	mongoDBClient, err := mongoDB.NewClientMongo(context)
	if err != nil {
		return
	}
	dbMongo := mongoDB.NewDb(mongoDBClient)
	tokenSigned, err := middleware.GetToken(context)
	if err != nil {
		errorRating := fmt.Errorf("get rating error : %v", err)
		dbMongo.AddNewLog(context, "-", errorRating, "-")
		HandleError(err, context)
		return
	}
	db, err := dataBase.NewDB(configs.NewConfig())
	if err != nil {
		errorRating := fmt.Errorf("get rating error : %v", err)
		dbMongo.AddNewLog(context, "-", errorRating, "-")
		HandleError(err, context)
		return
	}
	claims, err := services.VerifyToken(tokenSigned)
	if err != nil {
		errorRating := fmt.Errorf("get rating error : %v", err)
		dbMongo.AddNewLog(context, "-", errorRating, "-")
		HandleError(err, context)
		return
	}
	userPhone := claims.Phone
	rating, err := db.GetRatingByPhone(userPhone)
	if err != nil {
		errorRating := fmt.Errorf("get rating error : %v", err)
		dbMongo.AddNewLog(context, userPhone, errorRating, "-")
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"phone":  userPhone,
		"rating": rating,
	})
	dbMongo.AddNewLog(context, userPhone, nil, "")
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
	//TODO : ere will me method that address to order service`s database and insert rating to the last trip of user, if time is not expired
}
