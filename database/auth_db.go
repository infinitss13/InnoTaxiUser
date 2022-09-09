package database

import (
	"time"

	"github.com/infinitss13/innotaxiuser/entity"
	_ "github.com/lib/pq"
)

//InsertUser - function that will implement new user to the DB
func (dataBase *DataBase) InsertUser(user entity.User) error {
	timeNow := time.Now()
	query := "INSERT INTO users (name, phone, email, password_hash,rating, created_at, updated_at, deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	row := dataBase.db.QueryRow(query, user.Name, user.Phone, user.Email, user.Password, 0.0, timeNow, timeNow, false)
	var id = 0
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

//UserIsRegistered - function check the user in DB by phone number
func (dataBase *DataBase) UserIsRegistered(userPhone string) (string, error) {
	query := "SELECT password_hash FROM users WHERE phone=$1"
	var passwordHash string
	err := dataBase.db.Get(&passwordHash, query, userPhone)
	if err != nil {
		return "", UserNotFound
	}
	return passwordHash, nil
}

//UserExist - function check if the user already exists
func (dataBase *DataBase) UserExist(user entity.User) (bool, error) {
	query := "SELECT id FROM users WHERE phone=$1 OR email=$2"
	var id int
	err := dataBase.db.Get(&id, query, user.Phone, user.Email)
	if err != nil {
		return false, err
	}

	return true, nil

}
