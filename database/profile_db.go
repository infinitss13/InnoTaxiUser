package database

import (
	"github.com/infinitss13/innotaxiuser/entity"
)

func (dataBase *DB) GetUserByPhone(userPhone string) (entity.ProfileData, error) {
	query := "SELECT name, phone, email,rating from users WHERE phone=$1"
	user := entity.ProfileData{}
	err := dataBase.db.Get(&user, query, userPhone)
	if err != nil {
		return entity.ProfileData{}, err
	}
	return user, nil
}
