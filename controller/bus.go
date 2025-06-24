package controller

import (
	"bus-booking/constant"
	"bus-booking/entity"
	"bus-booking/repository"
	"bus-booking/usecase"
	"bus-booking/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

type BusController struct {
	Mysql *repository.MysqlCon
}

func (b *BusController) BusCreate(c echo.Context) error {
	verifyId := c.Get("userID")
	// fmt.Println("veridyId----", veridyId)
	userUsecase := usecase.UserUsecase{
		Mysql: b.Mysql,
	}

	user, err := userUsecase.FindUser(verifyId)
	if err != nil {
		return c.JSON(constant.BAD_REQUEST, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.BAD_REQUEST_MESSAGE,
			Error:   err,
		})
	}

	if user.UserType != constant.ADMIN {
		return c.JSON(constant.BAD_REQUEST, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.USER_ACCOUNT,
		})
	}

	req := entity.Bus{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(constant.BAD_REQUEST, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.BAD_REQUEST_MESSAGE,
		})
	}
	req.BusID = "BUS_" + utils.GenerateUserID()
	fmt.Println(req.BusID)
	busUsecase := usecase.BusUsecase{
		Mysql: b.Mysql,
	}

	if err := busUsecase.CreateBus(req); err != nil {
		return c.JSON(constant.BAD_REQUEST, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.BAD_REQUEST_MESSAGE,
		})
	}
	return c.JSON(constant.SUCCESS, entity.Response{
		Status:  constant.SUCCESS,
		Message: constant.SUCCESS_MESSAGE,
	})

}
