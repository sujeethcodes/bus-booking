package middleware

import (
	"bus-booking/constant"
	"bus-booking/entity"
	"bus-booking/repository"
	"bus-booking/usecase"

	"github.com/labstack/echo/v4"
)

func AdminMiddleware(db *repository.MysqlCon) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID, _ := c.Get("userID").(string)
			userUseCase := usecase.UserUsecase{Mysql: db}
			user, err := userUseCase.FindUser(userID)
			if err != nil || user.UserType != constant.ADMIN {
				return c.JSON(constant.BAD_REQUEST, entity.Response{
					Status:  constant.BAD_REQUEST,
					Message: constant.USER_ACCOUNT,
				})
			}
			return next(c)
		}
	}
}
