package controller

import (
	"bus-booking/constant"
	"bus-booking/entity"
	"bus-booking/middleware"
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
		})
	}
	userUsecase := usecase.UserUsecase{
		Mysql: u.Mysql,
	}
	details, _ := userUsecase.IsEmailExit(req.Email)
	fmt.Println("email-----", details.Email)
	if len(details.Email) != 0 {
		return c.JSON(constant.BAD_REQUEST, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.EMAIL_EXIST,
		})
	}
	if req.UserType != constant.ADMIN && req.UserType != constant.USER {
		return c.JSON(constant.BAD_REQUEST, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.INVAILD_USER_TYPE_MESSAGE,
		})
	}
	req.UserID = utils.GenerateUserID()
	if req.UserType == constant.ADMIN {
		req.UserID = "adm_" + req.UserID
	}
	token, err := middleware.GenerateToken(req.UserID)
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

func (u *UserController) EditUser(c echo.Context) error {
	user_id := c.QueryParam("user_id")
	name := c.QueryParam("name")
	email := c.QueryParam("email")
	password := c.QueryParam("password")
	phone_number := c.QueryParam("phone_number")
	address := c.QueryParam("address")
	status := c.QueryParam("status")

	if err := utils.VerifyUserId(user_id, c); err != nil {
		return c.JSON(constant.UNAUTHORIZED, entity.Response{
			Status:  constant.UNAUTHORIZED,
			Message: constant.INVAILD_USER_ID,
		})
	}
	req := entity.EditUserReq{
		UserID:      user_id,
		Name:        name,
		Email:       email,
		Password:    password,
		PhoneNumber: phone_number,
		Address:     json.RawMessage(address),
		Status:      status,
	}

	userUsecase := usecase.UserUsecase{
		Mysql: u.Mysql,
	}

	err := userUsecase.EditUser(&req)
	if err != nil {
		return c.JSON(constant.BAD_REQUEST, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.DB_USER_UPDATE_FAIL,
			Error:   err.Error(),
		})
	}
	return c.JSON(constant.SUCCESS, entity.Response{
		Status:  constant.SUCCESS,
		Message: constant.SUCCESS_MESSAGE,
	})
}

func (u *UserController) DeleteUser(c echo.Context) error {

	req := entity.DeleteUserReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(constant.BAD_REQUEST, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.BAD_REQUEST_MESSAGE,
		})
	}

	userUsecase := usecase.UserUsecase{
		Mysql: u.Mysql,
	}
	req.Status = constant.IN_ACTIVE
	err := userUsecase.DeleteUser(req)
	if err != nil {
		return c.JSON(constant.BAD_REQUEST, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.DB_USER_DELETE_FAIL,
			Error:   err.Error(),
		})
	}
	return c.JSON(constant.SUCCESS, entity.Response{
		Status:  constant.SUCCESS,
		Message: constant.DELETE_SUCCESS_MESSAGE,
	})
}
