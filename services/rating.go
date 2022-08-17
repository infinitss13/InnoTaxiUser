package services

func (srv *Service) GetRatingWithToken(tokenSigned string) (float32, string, error) {

	claims, err := VerifyToken(tokenSigned)
	if err != nil {
		return 0, "", err
	}
	userPhone := claims.Phone
	rating, err := srv.Db.GetRatingByPhone(userPhone)
	if err != nil {
		return 0, "", err
	}
	return rating, userPhone, nil
}
