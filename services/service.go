package services

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/infinitss13/innotaxiuser/configs"
	"github.com/infinitss13/innotaxiuser/database"
	"github.com/infinitss13/innotaxiuser/entity"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type UserService interface {
	CreateUser(entity.User) error
	SignInUser(entity.InputSignIn) (string, error)
	GetUserByToken(string) (entity.ProfileData, error)
	UpdateUserProfile(string, *entity.UpdateData) error
	DeleteProfile(string, password string) error
	GetRatingWithToken(string) (float32, string, error)
	VerifyToken(tokenSigned string) (entity.InputSignIn, error)
	GetToken(context *gin.Context) (string, error)
	Auth() gin.HandlerFunc
}

type Service struct {
	Db database.DataBase
}

func NewService(db database.DataBase) Service {
	return Service{Db: db}
}

func (srv Service) GetToken(context *gin.Context) (string, error) {
	tokenString := context.GetHeader("Authorization")
	if tokenString == "" {
		context.JSON(401, gin.H{"error": "request does not contain an access token"})
		context.Abort()
		return "", errors.New("no access token")
	}
	splitedToken := strings.Split(tokenString, " ")
	return splitedToken[1], nil
}

func (srv Service) Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		splitedToken, err := srv.GetToken(context)
		if err != nil {
			logrus.Error(err)
			context.JSON(http.StatusInternalServerError, err)
		}
		_, err = srv.VerifyToken(splitedToken)
		//_, err = services.VerifyToken(splitedToken)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
	}
}

func (srv Service) CreateUser(user entity.User) error {
	password, err := GenerateHash(user.Password)
	if err != nil {
		return err
	}
	user.Password = password
	isExist, _ := srv.Db.UserExist(user)
	if isExist {
		return database.UserExistErr
	}

	err = srv.Db.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (srv Service) SignInUser(inputSignIn entity.InputSignIn) (string, error) {
	passwordHash, err := srv.Db.UserIsRegistered(inputSignIn.Phone)
	if err != nil {
		return "", err
	}
	if err = CheckPassword(inputSignIn.Password, passwordHash); err != nil {
		return "", err
	}
	token, err := CreateToken(inputSignIn.Phone)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (srv Service) GetRatingWithToken(tokenSigned string) (float32, string, error) {
	claims, err := srv.VerifyToken(tokenSigned)
	if err != nil {
		return 0, "", err
	}
	userPhone := claims.Phone
	rating, err := srv.Db.GetRatingByPhone(userPhone)
	if err != nil {
		return 0, "", err
	}
	return rating, userPhone, nil
}

func (srv Service) GetUserByToken(tokenSigned string) (entity.ProfileData, error) {
	claims, err := srv.VerifyToken(tokenSigned)
	if err != nil {
		return entity.ProfileData{}, err
	}
	userPhone := claims.Phone
	user, err := srv.Db.GetUserByPhone(userPhone)
	if err != nil {
		return entity.ProfileData{}, err
	}
	return user, nil
}

func (srv Service) UpdateUserProfile(tokenSigned string, data *entity.UpdateData) error {
	claims, err := srv.VerifyToken(tokenSigned)
	if err != nil {
		return err
	}
	isCorrect, _ := srv.Db.CheckUpdateData(claims.Phone, data)
	if !isCorrect {
		return database.UpdateDataError
	}
	userPhone := claims.Phone
	err = srv.Db.UpdateUser(userPhone, data)
	if err != nil {
		return err
	}
	return nil
}

func (srv Service) DeleteProfile(tokenSigned string, password string) error {
	claims, err := srv.VerifyToken(tokenSigned)
	if err != nil {
		return err
	}
	passwordHash, err := srv.Db.UserIsRegistered(claims.Phone)
	if err != nil {
		return err
	}
	if err = CheckPassword(password, passwordHash); err != nil {
		return err
	}
	err = srv.Db.DeleteProfile(claims.Phone)
	if err != nil {
		return err
	}
	return nil

}

func GenerateHash(password string) (string, error) {
	saltedBytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashPassword := string(hashedBytes)
	return hashPassword, nil
}
func CheckPassword(password, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err
}

type JWTClaim struct {
	Phone string `json:"phone"`
	jwt.StandardClaims
}

func CreateToken(phone string) (string, error) {
	timeExp, err := strconv.Atoi(configs.GetEnv("TOKEN_EXPIRES", "15"))
	if err != nil {
		return "", err
	}
	claims := &JWTClaim{
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(timeExp) * time.Minute).Unix(),
		},
	}
	fmt.Println(claims.ExpiresAt)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (srv Service) VerifyToken(tokenSigned string) (entity.InputSignIn, error) {
	signInData := entity.InputSignIn{}
	token, err := jwt.ParseWithClaims(tokenSigned, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_KEY")), nil
	})

	if err != nil {
		return signInData, errors.New("Error parsing claims")
	}
	claims := token.Claims.(*JWTClaim)
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return signInData, errors.New("token expired")
	}
	signInData.Phone = claims.Phone
	return signInData, nil
}
