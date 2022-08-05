package dataBase

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/infinitss13/InnoTaxiUser"
	"github.com/infinitss13/InnoTaxiUser/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//InsertUser - function that will implement new user to the DB
func InsertUser(user *InnoTaxiUser.User) (int, error) {
	var envVariables = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("HOST_DB"),
		os.Getenv("PORT_DB"),
		os.Getenv("USERNAME_DB"),
		os.Getenv("DBNAME_DB"),
		os.Getenv("PASSWORD_DB"),
		os.Getenv("SSLMODE_DB"))

	db, err := sqlx.Open("postgres", envVariables)

	if err != nil {
		return 0, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
		return 0, err
	}
	id, err := UserExist(user)
	if id != 0 {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (name, phone, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id", "users")
	row := db.QueryRow(query, user.Name, user.Phone, user.Email, user.Password)
	id = 0
	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil

}

//CheckUser - function check the user data in DB
func CheckUser(userPhone, password string) error {
	var envVariables = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("HOST_DB"),
		os.Getenv("PORT_DB"),
		os.Getenv("USERNAME_DB"),
		os.Getenv("DBNAME_DB"),
		os.Getenv("PASSWORD_DB"),
		os.Getenv("SSLMODE_DB"))
	db, err := sqlx.Connect("postgres", envVariables)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("SELECT password_hash FROM users WHERE phone=$1")
	var password_hash string
	err = db.Get(&password_hash, query, userPhone)
	if err != nil {
		return err
	}
	err = services.CheckPassword(password, password_hash)
	if err != nil {
		return err
	}
	return nil
}

//UserExist - function check if the user already exists and can't sign-up
func UserExist(user *InnoTaxiUser.User) (int, error) {
	var envVariables = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("HOST_DB"),
		os.Getenv("PORT_DB"),
		os.Getenv("USERNAME_DB"),
		os.Getenv("DBNAME_DB"),
		os.Getenv("PASSWORD_DB"),
		os.Getenv("SSLMODE_DB"))
	db, err := sqlx.Connect("postgres", envVariables)
	if err != nil {
		return -1, err
	}
	err = db.Ping()
	if err != nil {
		return -1, err
	}
	query := fmt.Sprintf("SELECT id FROM users WHERE phone=$1 OR email=$2")
	var id int
	err = db.Get(&id, query, user.Phone, user.Email)
	if err != nil {
		return 0, nil
	}
	fmt.Println(err)
	fmt.Printf("Here is id of user %d: ", id)
	return id, errors.New("user exists")

}
