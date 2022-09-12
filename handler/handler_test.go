package handler

import (
	"bytes"
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
			token:                "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZSI6IiszNzU0NDM0NzIzNDIiLCJleHAiOjE2OTQ1MDM3NzR9.8czPPLxHSJXJ9qBHZggjNOx6PlJr1dp6SQ6Vv7h-XvU",
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: "hello",
		},
	}
	w := httptest.NewRecorder()
	r := gin.New()
	r.GET("/api/profile", handler.getProfile)

	cache.EXPECT().GetValue(tests[0].token).Return(false, nil)
	service.EXPECT().GetToken(gomock.Any()).Return(tests[0].token, nil)
	service.EXPECT().GetUserByToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZSI6IiszNzU0NDM0NzIzNDIiLCJleHAiOjE2OTQ1MDM3NzR9.8czPPLxHSJXJ9qBHZggjNOx6PlJr1dp6SQ6Vv7h-XvU").Return(userInfo, nil)
	responseBody := `{"Name":"polina","Phone":"+375443472342","Email":"polina@gmail.com","Rating":0}`
	req := httptest.NewRequest("GET", "/api/profile", nil)

	r.ServeHTTP(w, req)
	t.Log(w.Body)
	t.Log(responseBody)
	assert.Equal(t, w.Body.String(), responseBody)
	assert.Equal(t, w.Code, http.StatusOK)

}
