package dataBase

import (
	"fmt"
)

func (dataBase *DB) GetRatingByPhone(userPhone string) (float32, error) {
	query := fmt.Sprintf("SELECT rating from %s WHERE phone=$1", "users")
	var rating float32
	err := dataBase.db.Get(&rating, query, userPhone)
	if err != nil {
		return 0, err
	}
	return rating, nil
}
