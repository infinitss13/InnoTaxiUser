package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/InnoTaxiUser/configs"
	"github.com/infinitss13/InnoTaxiUser/dataBase"
	"github.com/infinitss13/InnoTaxiUser/dataBase/mongoDB"
	"github.com/infinitss13/InnoTaxiUser/entity"
	"github.com/infinitss13/InnoTaxiUser/services"
	"github.com/sirupsen/logrus"
)

func signUp(context *gin.Context) {
	mongoDBClient, err := mongoDB.NewClientMongo(context)
	if err != nil {
		return
	}
	dbMongo := mongoDB.NewDb(mongoDBClient)
	input := new(entity.User)
	err = context.BindJSON(&input)
	if err != nil {
		dbMongo.AddNewLog(context, input.Phone, err, "some problems")
		HandleError(err, context)
		return
	}
	input, err = services.CreateUser(input)
	if err != nil {
		errorCreate := fmt.Errorf("SIGN-UP ERROR, %v", err)
		dbMongo.AddNewLog(context, input.Phone, errorCreate, "some problems")
		HandleError(err, context)
		return
	}

	db, err := dataBase.NewDB(configs.NewConfig())
	id, err := db.InsertUser(*input)

	if err != nil {
		errorCreate := fmt.Errorf("sign-up error,%v", err)
		dbMongo.AddNewLog(context, input.Phone, errorCreate, "some problems")
		HandleError(err, context)
		return
	}
	err = dbMongo.AddNewLog(context, input.Phone, nil, "")
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, id)
	logrus.Info("status code: ", http.StatusOK, " User is created")
}

func signIn(context *gin.Context) {
	input := struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	mongoDBClient, errorMongo := mongoDB.NewClientMongo(context)
	if errorMongo != nil {
		return
	}
	dbMongo := mongoDB.NewDb(mongoDBClient)

	if err := context.BindJSON(&input); err != nil {
		errorCreate := fmt.Errorf("sign-in error,%v", err)
		dbMongo.AddNewLog(context, input.Phone, errorCreate, "some problems")
		ErrorBinding(context)
		return
	}
	db, err := dataBase.NewDB(configs.NewConfig())
	if err != nil {
		errorCreate := fmt.Errorf("postgres error,%v", err)
		dbMongo.AddNewLog(context, input.Phone, errorCreate, "some problems")
		HandleError(err, context)
		return
	}
	err = db.UserIsRegistered(input.Phone, input.Password)
	if err != nil {
		errorCreate := fmt.Errorf("sign-in error,%v", err)
		dbMongo.AddNewLog(context, input.Phone, errorCreate, "some problems")
		HandleError(err, context)
		return
	}
	token, err := services.CreateToken(input.Phone)
	if err != nil {
		errorCreate := fmt.Errorf("sign-in error,%v", err)
		dbMongo.AddNewLog(context, input.Phone, errorCreate, "some problems")
		HandleError(err, context)
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello",
		"token":   token,
		"time":    time.Now(),
	})
	dbMongo.AddNewLog(context, input.Phone, nil, "")
	logrus.Info("status code :", http.StatusOK, " user is authorized")

}
