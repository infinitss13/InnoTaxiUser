package database

import (
	"github.com/infinitss13/innotaxiuser/entity"
)

func (dataBase *DataBase) GetUserByPhone(userPhone string) (entity.ProfileData, error) {
	query := "SELECT name, phone, email,rating from users WHERE phone=$1"
	user := entity.ProfileData{}
	err := dataBase.db.Get(&user, query, userPhone)
	if err != nil {
		return entity.ProfileData{}, err
	}
	return user, nil
}

func (dataBase *DataBase) UpdateUser(userPhone string, data *entity.UpdateData) error {
	query := "UPDATE users SET name=$1, phone=$2, email=$3 WHERE phone=$4"

	row := dataBase.db.QueryRow(query, data.Name, data.Phone, data.Email, userPhone)
	if row.Err() != nil {
		return UpdateProfileErr
	}
	return nil
}

func (dataBase *DataBase) CheckUpdateData(phone string, data *entity.UpdateData) (bool, error) {
	query := "SELECT id FROM users WHERE phone<>$1 AND(phone=$2 OR email=$3)"
	var id int
	err := dataBase.db.Get(&id, query, phone, data.Phone, data.Email)
	if err != nil {
		return true, err
	}

	return false, nil
}

func (dataBase *DataBase) DeleteProfile(phone string) error {
	query := "UPDATE users SET deleted=true WHERE phone=$1"

	row := dataBase.db.QueryRow(query, phone)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}
