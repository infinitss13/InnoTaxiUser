package dataBase

import (
	"github.com/infinitss13/InnoTaxiUser/configs"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

func NewDB(config configs.DBConfig) (*DB, error) {
	dataBase, err := sqlx.Open("postgres", config.ConnectionDbData())
	if err != nil {
		return nil, err
	}
	err = dataBase.Ping()
	if err != nil {
		return nil, err
	}
	return &DB{
		db: dataBase,
	}, nil
}
