package controller

// import (
// 	"bus-booking/constant"
// 	"bus-booking/entity"

// 	"github.com/labstack/echo/v4"
// )

// func CreateUser(c echo.Context) error {
// 	req := entity.User{}
// 	if err := c.Bind(&req); err != nil {
// 		return c.JSON(constant.BAD_REQUEST, entity.Response{
// 			Status:  constant.BAD_REQUEST,
// 			Message: constant.BAD_REQUEST_MESSAGE,
// 			Error:   err.Error(),
// 		})
// 	}
// }
