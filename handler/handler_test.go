package handler

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/infinitss13/innotaxiuser/entity"
	"github.com/infinitss13/innotaxiuser/mock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type LoggerTest struct {
}

func NewLoggerTest() LoggerTest {
	return LoggerTest{}
}

func (l LoggerTest) LogInfo(ctx *gin.Context) error {
	logrus.Info(ctx.Request.Method)
	return nil
}
func (l LoggerTest) LogError(ctx *gin.Context, err error) error {
	logrus.Info(ctx.Request.Method, err)
	return nil
}

func TestHandler_signUp(t *testing.T) {
	userStruct := entity.User{
		Name:     "stas",
		Phone:    "+375298913459",
		Email:    "stasrus23s@gmail.com",
		Password: "qwerty",
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	log := NewLoggerTest()

	service := mock.NewMockUserService(ctl)
	cache := mock.NewMockCache(ctl)
	handler := NewAuthHandlers(log, service, cache)

	r := gin.New()
	service.EXPECT().CreateUser(userStruct).Return(nil)
	r.POST("/sign-up", handler.signUp)
	w := httptest.NewRecorder()
	inputBody := `{"Name":"stas", "Phone":"+375298913459", "Email":"stasrus23s@gmail.com", "Password": "qwerty"}`

	expectedResponseBody := `"user successfully created"`
	req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(inputBody))
	r.ServeHTTP(w, req)
	t.Log(w.Body)
	assert.Equal(t, w.Body.String(), expectedResponseBody)
	assert.Equal(t, w.Code, http.StatusOK)
}

func TestHandler_getProfile(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	log := NewLoggerTest()
	service := mock.NewMockUserService(ctl)
	cache := mock.NewMockCache(ctl)
	handler := NewAuthHandlers(log, service, cache)
	userInfo := entity.ProfileData{
		Name:   "polina",
		Phone:  "+375443472342",
		Email:  "polina@gmail.com",
		Rating: 0,
	}

	tests := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "OK",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZSI6IiszNzU0NDM0NzIzNDIiLCJleHAiOjE2OTQyNDExMDZ9.GV4-JCHNYqIP8wVUsxK3wlqaLk62xzsgzRtA3sTXOJE",
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: "hello",
		},
	}

	service.EXPECT().GetToken(context.Context()).Return(tests[0].token, nil)
	service.EXPECT().GetUserByToken(tests[0].token).Return(userInfo, nil)
	r := gin.New()
	r.GET("/api/profile", handler.getProfile)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/profile", nil)
	req.Header.Set(tests[0].headerName, tests[0].headerValue+tests[0].token)
	r.ServeHTTP(w, req)
	t.Log(w.Body)
	assert.Equal(t, w.Code, http.StatusOK)

}
