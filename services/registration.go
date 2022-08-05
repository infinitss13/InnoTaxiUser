package services

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/infinitss13/InnoTaxiUser/entity"
	"golang.org/x/crypto/bcrypt"
)

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
