package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/innotaxiuser/entity"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

// @Summary Get Profile
// @Security ApiKeyAuth
// @Tags profile
// @Description handler for get profile request, allows user to check his profile data(name,  phone number, email, rating)
// @ID getprofile
// @Param input body entity.User true "account info"
// @Produce json
// @Success 200 {object} entity.ProfileData
// @Failure 400 {object} error
// @Router /api/profile [get]
func (handler AuthHandlers) getProfile(context *gin.Context) {
	requestGetProfile.Inc()
	requestProcessed.Inc()
	timer := prometheus.NewTimer(httpDuration.WithLabelValues(context.Request.RequestURI))

	tokenSigned, err := handler.GetAndCheckToken(context)
	if err != nil {
		errorToken := fmt.Errorf("profile error,%v", err)
		if errorLogger := handler.LoggerMongo.LogError(context, errorToken); errorLogger != nil {
			logrus.Error("error in logger", errorLogger)
		}
		HandleError(err, context)
		return
	}
	userData, err := handler.UserService.GetUserByToken(tokenSigned)

	if err != nil {
		errorToken := fmt.Errorf("profile error,%v", err)
		if errorLogger := handler.LoggerMongo.LogError(context, errorToken); errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger.Error())
		}
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, userData)
	if err = handler.LoggerMongo.LogInfo(context); err != nil {
		logrus.Error("error in logger : ", err)
	}
	logrus.Info("status code :", http.StatusOK, " user get profile")
	timer.ObserveDuration()

}

// @Summary Update Profile
// @Security ApiKeyAuth
// @Tags profile
// @Description handler for get profile request, allows user to update data(name, phone number, email). You can change any field of this data, but you should input correct data"
// @ID updateprofile
// @Accept json
// @Produce json
// @Param input body entity.UpdateData true "user's new data"
// @Success 200 {object} string "user updated profile"
// @Failure 400 {object} error
// @Router /api/profile [patch]
func (handler AuthHandlers) updateProfile(context *gin.Context) {
	update := new(entity.UpdateData)
	if err := context.BindJSON(&update); err != nil {
		errorCreate := fmt.Errorf("update profile error: %v", err)
		if errorLogger := handler.LoggerMongo.LogError(context, errorCreate); errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger)
		}
		ErrorBinding(context)
		return
	}

	tokenSigned, err := handler.GetAndCheckToken(context)
	if err != nil {
		errorToken := fmt.Errorf("update profile error,%v", err)
		if errorLogger := handler.LoggerMongo.LogError(context, errorToken); errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger.Error())
		}
		HandleError(err, context)
		return
	}

	err = handler.UserService.UpdateUserProfile(tokenSigned, update)
	if err != nil {
		errorUpdate := fmt.Errorf("update profile error,%v", err)
		if errorLogger := handler.LoggerMongo.LogError(context, errorUpdate); errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger.Error())
		}
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, "User is successfully updated")
	logrus.Info("status code :", http.StatusOK, " user updated profile")
	if err = handler.Cache.SetValue(tokenSigned, "true"); err != nil {
		logrus.Error("error in cache db : ", err.Error())
		if errorLogger := handler.LoggerMongo.LogError(context, err); errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger.Error())
		}
	}
	if err = handler.LoggerMongo.LogInfo(context); err != nil {
		logrus.Error("error in logger : ", err)
	}
}

// @Summary Delete Profile
// @Security ApiKeyAuth
// @Tags profile
// @Description handler for delete profile request, allows user to delete his account(user's data will be available for registration for other users)
// @ID deleteprofile
// @Param input body entity.DeleteData true "password"
// @Accept json
// @Produce json
// @Success 200 {object} string "User is successfully deleted"
// @Failure 400 {object} error
// @Router /api/profile [delete]
func (handler AuthHandlers) deleteProfile(context *gin.Context) {
	deleteData := new(entity.DeleteData)
	if err := context.BindJSON(&deleteData); err != nil {
		errorCreate := fmt.Errorf("delete profile error: %v", err)
		if errorLogger := handler.LoggerMongo.LogError(context, errorCreate); errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger.Error())
		}
		ErrorBinding(context)
		return
	}
	tokenSigned, err := handler.GetAndCheckToken(context)
	if err != nil {
		errorToken := fmt.Errorf("delete profile error: %v", err)
		if errorLogger := handler.LoggerMongo.LogError(context, errorToken); errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger.Error())
		}
		HandleError(err, context)
		return
	}
	err = handler.UserService.DeleteProfile(tokenSigned, deleteData.Password)
	if err != nil {
		errorToken := fmt.Errorf("delete profile error: %v", err)
		errorLog := handler.LoggerMongo.LogError(context, errorToken)
		if errorLog != nil {
			logrus.Info("error in logger : ", errorLog.Error())
		}
		HandleError(err, context)
		return
	}
	context.JSON(http.StatusOK, "User is successfully deleted")
	logrus.Info("status code :", http.StatusOK, " user deleted profile")
	if err = handler.Cache.SetValue(tokenSigned, "true"); err != nil {
		logrus.Error("error in cache db : ", err.Error())
		if errorLogger := handler.LoggerMongo.LogError(context, err); errorLogger != nil {
			logrus.Error("error in logger : ", errorLogger.Error())
		}
	}
	if err = handler.LoggerMongo.LogInfo(context); err != nil {
		logrus.Error("error in logger : ", err.Error())
	}
}
