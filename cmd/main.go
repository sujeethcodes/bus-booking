package main

import (
	"bus-booking/connectors"
	"bus-booking/controller"
	"bus-booking/middleware"
	"bus-booking/repository"

	"os"

	"github.com/labstack/echo/v4"
)

type Container struct {
	UserInstance controller.UserController
	BusInstance  controller.BusController
}

func LoadContainer() *Container {
	return &Container{
		UserInstance: controller.UserController{Mysql: repository.SingletonMysqlCon()},
		BusInstance:  controller.BusController{Mysql: repository.SingletonMysqlCon()},
	}
}

func init() {
	connectors.LoadEnv()
}

func main() {
	containerInstane := LoadContainer()
	e := echo.New()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// user protect group
	userGropu := e.Group("/user", middleware.JWTMiddleware())
	// USER - Routes
	e.POST("/user", containerInstane.UserInstance.CreateUser)
	userGropu.PUT("", containerInstane.UserInstance.EditUser)
	userGropu.DELETE("", containerInstane.UserInstance.DeleteUser)
	e.POST("/bus", containerInstane.BusInstance.BusCreate)

	e.Start(":" + PORT)
}
