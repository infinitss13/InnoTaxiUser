package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/infinitss13/innotaxiuser/configs"
	"github.com/infinitss13/innotaxiuser/database"
)

type OrderData struct {
	TaxiType string `json:"taxiType"`
	From     string `json:"from"`
	To       string `json:"to"`
}

func orderTaxi(context *gin.Context) {
	input := new(OrderData)
	err := context.BindJSON(&input)
	if err != nil {
		ErrorBinding(context)
		return
	}
	_, err = database.NewDB(configs.NewConfig())
	if err != nil {
		HandleError(err, context)
		return
	}

	//TODO : Here should be method that will address the Order service
}
