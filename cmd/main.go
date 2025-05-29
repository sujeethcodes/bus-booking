package main

import (
	"bus-booking/connectors"
	"bus-booking/controller"
	"bus-booking/repository"

	"os"

	"github.com/labstack/echo/v4"
)

type Container struct {
	UserInstance controller.UserController
}

func LoadContainer() *Container {
	return &Container{
		UserInstance: controller.UserController{Mysql: repository.SingletonMysqlCon()},
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
	// USER - Routes
	e.POST("/user", containerInstane.UserInstance.CreateUser)

	e.Start(":" + PORT)
}
