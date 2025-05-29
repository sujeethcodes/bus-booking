package controller

import (
	"bus-booking/constant"
	"bus-booking/entity"
	"bus-booking/repository"
	"bus-booking/usecase"
	"bus-booking/utils"
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Mysql *repository.MysqlCon
}

func (u *UserController) CreateUser(c echo.Context) error {
	req := entity.User{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(constant.BAD_REQUEST, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.BAD_REQUEST_MESSAGE,
			Error:   err.Error(),
		})
	}
	req.UserID = utils.GenerateUserID()
	token, err := utils.GenerateToken(req.UserID)
	fmt.Println("token----", token)
	fmt.Println("len token----", len(token))

	req.Status = constant.ACTIVE
	if err != nil {
		return c.JSON(constant.BAD_REQUEST, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.TOKEN_GENERATE_FAILED,
			Error:   err.Error(),
		})
	}
	req.Token = token
	userUsecase := usecase.UserUsecase{
		Mysql: u.Mysql,
	}
	jsonData, _ := json.Marshal(req.Address)
	req.Address = jsonData

	if err := userUsecase.CreateUser(req); err != nil {
		return c.JSON(constant.BAD_REQUEST, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.BAD_REQUEST_MESSAGE,
			Error:   err.Error(),
		})
	}
	return c.JSON(constant.SUCCESS, entity.Response{
		Status:  constant.SUCCESS,
		Message: constant.SUCCESS_MESSAGE,
	})
}
