package dataBase

import (
	"fmt"
	"log"
	"time"

	"github.com/infinitss13/InnoTaxiUser/configs"
	"github.com/infinitss13/InnoTaxiUser/entity"
	"github.com/infinitss13/InnoTaxiUser/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//InsertUser - function that will implement new user to the DB
func InsertUser(config *configs.DBConfig, user *entity.User) (int, error) {

	db, err := sqlx.Open("postgres", config.ConnectionDbData())

	if err != nil {
		return 0, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
		return 0, err
	}
	id, err := UserExist(config, user)
	if id != 0 {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (name, phone, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", "users")
	row := db.QueryRow(query, user.Name, user.Phone, user.Email, user.Password, time.Now(), time.Now())
	id = 0
	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil

}

//CheckUser - function check the user data in DB
func CheckUser(config *configs.DBConfig, userPhone, password string) error {
	db, err := sqlx.Connect("postgres", config.ConnectionDbData())
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("SELECT password_hash FROM users WHERE phone=$1")
	var passwordHash string
	err = db.Get(&passwordHash, query, userPhone)
	if err != nil {
		return err
	}
	err = services.CheckPassword(password, passwordHash)
	if err != nil {
		return err
	}
	return nil
}

//UserExist - function check if the user already exists and can't sign up
func UserExist(config *configs.DBConfig, user *entity.User) (int, error) {
	db, err := sqlx.Connect("postgres", config.ConnectionDbData())
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

	return id, UserExistErr

}
