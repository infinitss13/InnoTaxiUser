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

// @Summary Order Taxi
// @Security ApiKeyAuth
// @Tags order
// @Description handler for get order request, allows user to order the taxi
// @ID ordertaxi
// @Param input body OrderData true "order info"
// @Produce json
// @Success 200 {object} string "all is okay"
// @Failure 500 {object} error
// @Router /api/order [post]
func orderTaxi(context *gin.Context) {
	input := new(OrderData)
	err := context.BindJSON(&input)
	if err != nil {
		ErrorBinding(context)
		return
	}
	_, err = database.NewDataBase(configs.NewDBConfig())
	if err != nil {
		HandleError(err, context)
		return
	}

	//TODO : Here should be method that will address the Order service
}
