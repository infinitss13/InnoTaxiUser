package database

import (
	"github.com/infinitss13/innotaxiuser/configs"
	"github.com/infinitss13/innotaxiuser/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type DataBase struct {
	*sqlx.DB
}

func NewDataBase(config configs.DBConfig) (DataBase, error) {
	dataBase, err := sqlx.Open("postgres", config.ConnectionDbData())
	if err != nil {
		return DataBase{}, err
	}
	err = dataBase.Ping()
	if err != nil {
		return DataBase{}, err
	}
	return DataBase{
		dataBase,
	}, nil
}

//InsertUser - function that will implement new user to the DB
func (dataBase DataBase) InsertUser(user entity.User) error {
	timeNow := time.Now()
	query := "INSERT INTO users (name, phone, email, password_hash,rating, created_at, updated_at, deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	row := dataBase.QueryRow(query, user.Name, user.Phone, user.Email, user.Password, 0.0, timeNow, timeNow, false)
	var id = 0
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

//UserIsRegistered - function check the user in DB by phone number
func (dataBase DataBase) UserIsRegistered(userPhone string) (string, error) {
	query := "SELECT password_hash FROM users WHERE phone=$1"
	var passwordHash string
	err := dataBase.Get(&passwordHash, query, userPhone)
	if err != nil {
		return "", UserNotFound
	}
	return passwordHash, nil
}

//UserExist - function check if the user already exists
func (dataBase DataBase) UserExist(user entity.User) (bool, error) {
	query := "SELECT id FROM users WHERE phone=$1 OR email=$2"
	var id int
	err := dataBase.Get(&id, query, user.Phone, user.Email)
	if err != nil {
		return false, err
	}

	return true, nil

}

func (dataBase DataBase) GetUserByPhone(userPhone string) (entity.ProfileData, error) {
	query := "SELECT name, phone, email,rating from users WHERE phone=$1"
	user := entity.ProfileData{}
	err := dataBase.Get(&user, query, userPhone)
	if err != nil {
		return entity.ProfileData{}, err
	}
	return user, nil
}

func (dataBase DataBase) UpdateUser(userPhone string, data *entity.UpdateData) error {
	query := "UPDATE users SET name=$1, phone=$2, email=$3 WHERE phone=$4"
	row := dataBase.QueryRow(query, data.Name, data.Phone, data.Email, userPhone)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (dataBase DataBase) CheckUpdateDataAlreadyTaken(phone string, data *entity.UpdateData) (bool, error) {
	query := "SELECT id FROM users WHERE phone<>$1 AND(phone=$2 OR email=$3)"
	var id int
	err := dataBase.Get(&id, query, phone, data.Phone, data.Email)
	if err != nil {
		return true, err
	}
	return false, nil
}

func (dataBase DataBase) DeleteProfile(phone string) error {
	query := "UPDATE users SET deleted=true WHERE phone=$1"
	row := dataBase.QueryRow(query, phone)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (dataBase DataBase) GetRatingByPhone(userPhone string) (float32, error) {
	query := "SELECT rating from users WHERE phone=$1"
	var rating float32
	err := dataBase.Get(&rating, query, userPhone)
	if err != nil {
		return 0, err
	}
	return rating, nil
}
