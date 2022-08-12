package dataBase

func (dataBase *DB) GetRatingByPhone(userPhone string) (float32, error) {
	query := "SELECT rating from users WHERE phone=$1"
	var rating float32
	err := dataBase.db.Get(&rating, query, userPhone)
	if err != nil {
		return 0, err
	}
	return rating, nil
}
