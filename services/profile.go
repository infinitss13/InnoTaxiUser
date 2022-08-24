package services

import (
	"github.com/infinitss13/innotaxiuser/database"
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

func (srv *Service) UpdateUserProfile(tokenSigned string, data *entity.UpdateData) error {
	claims, err := VerifyToken(tokenSigned)
	if err != nil {
		return err
	}
	isCorrect, err := srv.Db.CheckUpdateData(claims.Phone, data)
	if isCorrect != true {
		return database.UpdateDataError
	}
	userPhone := claims.Phone
	err = srv.Db.UpdateUser(userPhone, data)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) DeleteProfile(tokenSigned string, password string) error {
	claims, err := VerifyToken(tokenSigned)
	if err != nil {
		return err
	}
	passwordHash, err := srv.Db.UserIsRegistered(claims.Phone)
	if err != nil {
		return err
	}
	if err = CheckPassword(password, passwordHash); err != nil {
		return err
	}
	err = srv.Db.DeleteProfile(claims.Phone)
	if err != nil {
		return err
	}
	return nil

}
