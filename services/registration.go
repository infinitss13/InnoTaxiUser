package services

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/infinitss13/InnoTaxiUser"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService struct {
}

func CreateUser(user *InnoTaxiUser.User) (*InnoTaxiUser.User, error) {
	user.Password = GenerateHash(user)

	//method that will insert the structure User to the db
	return user, nil
}

func GenerateHash(user *InnoTaxiUser.User) string {
	saltedBytes := []byte(user.Password)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	hashPassword := string(hashedBytes)
	return hashPassword

}
func CheckPassword(password, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err
}

func CreateToken(phone string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_phone"] = phone
	atClaims["exp_date"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_KEY")))
	if err != nil {
		return "", errors.New("troubles with jwt")
	}
	return token, nil
}
