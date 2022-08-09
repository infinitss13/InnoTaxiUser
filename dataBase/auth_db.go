package dataBase

import (
	"fmt"
	"time"

	"github.com/infinitss13/InnoTaxiUser/entity"
	"github.com/infinitss13/InnoTaxiUser/services"
	_ "github.com/lib/pq"
)

//InsertUser - function that will implement new user to the DB
func (dataBase *DB) InsertUser(user *entity.User) (int, error) {
	id, err := dataBase.UserExist(user)
	if id != 0 {
		return 0, err
	}
	timeNow := time.Now()
	query := fmt.Sprintf("INSERT INTO %s (name, phone, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", "users")
	row := dataBase.db.QueryRow(query, user.Name, user.Phone, user.Email, user.Password, timeNow, timeNow)
	id = 0
	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil

}

//UserIsRegistered - function check the user data in DB
func (dataBase *DB) UserIsRegistered(userPhone, password string) error {

	query := fmt.Sprintf("SELECT password_hash FROM users WHERE phone=$1")
	var passwordHash string
	err := dataBase.db.Get(&passwordHash, query, userPhone)
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
func (dataBase *DB) UserExist(user *entity.User) (int, error) {
	query := fmt.Sprintf("SELECT id FROM users WHERE phone=$1 OR email=$2")
	var id int
	err := dataBase.db.Get(&id, query, user.Phone, user.Email)
	if err != nil {
		return 0, nil
	}
	return id, UserExistErr

}
