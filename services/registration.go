package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/infinitss13/InnoTaxiUser/entity"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type SignInData struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func CreateUser(user *entity.User) (*entity.User, error) {
	password, err := GenerateHash(user)
	if err != nil {
		return nil, err
	}
	user.Password = password
	return user, nil
}

func GenerateHash(user *entity.User) (string, error) {
	saltedBytes := []byte(user.Password)
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

func VerifyToken(tokenSigned string) (SignInData, error) {

	signInData := SignInData{}
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
