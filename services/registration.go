package services

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/infinitss13/innotaxiuser/database"
	"github.com/infinitss13/innotaxiuser/entity"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Db *database.DB
}

func (srv *Service) CreateUser(user entity.User) error {
	password, err := GenerateHash(user.Password)
	if err != nil {
		return err
	}
	user.Password = password
	isExist, err := srv.Db.UserExist(user)
	if isExist != false || err != database.UserExistErr {
		return err
	}
	err = srv.Db.InsertUser(user)
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
	claims := &JWTClaim{
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenSigned string) (entity.InputSignIn, error) {

	signInData := entity.InputSignIn{}
	token, err := jwt.ParseWithClaims(tokenSigned, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_KEY")), nil
	})
	if err != nil {
		return signInData, err
	}
	claims := token.Claims.(*JWTClaim)
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return signInData, errors.New("token expired")
	}
	signInData.Phone = claims.Phone
	return signInData, nil
}

func (srv *Service) SignInUser(inputSignIn entity.InputSignIn) (string, error) {
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
