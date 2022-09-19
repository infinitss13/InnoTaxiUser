package database

import (
	"github.com/infinitss13/innotaxiuser/configs"
	"github.com/jmoiron/sqlx"
)

type DataBase struct {
	db *sqlx.DB
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
		db: dataBase,
	}, nil
}
