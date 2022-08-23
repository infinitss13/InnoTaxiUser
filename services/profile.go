package services

import (
	"github.com/infinitss13/innotaxiuser/entity"
)

func (srv *Service) GetUserByToken(tokenSigned string) (entity.ProfileData, error) {
	claims, err := VerifyToken(tokenSigned)
	if err != nil {
		return entity.ProfileData{}, err
	}
	userPhone := claims.Phone
	user, err := srv.Db.GetUserByPhone(userPhone)
	if err != nil {
		return entity.ProfileData{}, err
	}
	return user, nil
}
