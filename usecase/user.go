package usecase

import (
	"bus-booking/constant"
	"bus-booking/entity"
	"bus-booking/repository"
	"errors"
	"fmt"

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

func (u *UserUsecase) IsEmailExit(email string) (entity.User, error) {
	if u.Mysql == nil {
		zap.L().Info("Database connection failed")
		return entity.User{}, errors.New("database connection not initialized")
	}
	User := entity.User{}

	err := u.Mysql.Connection.Table(constant.USER_TABLE_NAME).Where("email = ?", email).First(&User).Error
	if err != nil {
		return User, err
	}
	return User, nil
}

func (u *UserUsecase) EditUser(req *entity.EditUserReq) error {
	fmt.Println("req---", req)
	if u.Mysql == nil {
		zap.L().Info("Database connection failed")
		return errors.New("database connection not initialized")
	}
	updates := make(map[string]interface{})

	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Password != "" {
		updates["password"] = req.Password
	}
	if req.PhoneNumber != "" {
		updates["phone_number"] = req.PhoneNumber
	}
	if len(req.Address) > 0 && string(req.Address) != "null" {
		updates["address"] = req.Address
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}

	// If no updates, return early
	if len(updates) == 0 {
		return errors.New("no fields provided to update")
	}
	err := u.Mysql.Connection.Table(constant.USER_TABLE_NAME).Where("user_id = ?", req.UserID).Updates(updates)
	if err != nil {
		return err.Error
	}
	return nil
}

func (u *UserUsecase) DeleteUser(req entity.DeleteUserReq) error {
	fmt.Println("req---", req)
	if u.Mysql == nil {
		zap.L().Info("Database connection failed")
		return errors.New("database connection not initialized")
	}

	err := u.Mysql.Connection.Table(constant.USER_TABLE_NAME).Where("user_id = ?", req.UserID).Updates(req).Error
	if err != nil {
		return errors.New("user delete failed")
	}
	return nil
}

func (u *UserUsecase) FindUser(userId interface{}) (entity.User, error) {
	fmt.Println("Enter find user")

	fmt.Println("userId--->", userId)
	user := entity.User{}
	if u.Mysql == nil {
		zap.L().Info("Database connection failed")
		return user, errors.New("database connection not initialized")
	}

	err := u.Mysql.Connection.Table(constant.USER).Where("user_id = ?", userId).Scan(&user).Error
	if err != nil {
		return user, errors.New("user fetch failed")
	}
	return user, nil
}
