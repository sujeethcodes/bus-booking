package usecase

import (
	"bus-booking/constant"
	"bus-booking/entity"
	"bus-booking/repository"
	"errors"

	"go.uber.org/zap"
)

type UserUsecase struct {
	Mysql *repository.MysqlCon
}

func (u *UserUsecase) CreateUser(req entity.User) error {
	if u.Mysql == nil {
		zap.L().Info("Database connection failed")
		return errors.New("database connection not initialized")
	}
	err := u.Mysql.Connection.Table(constant.USER_TABLE_NAME).Create(&req).Error
	if err != nil {
		return err
	}
	return nil

}
