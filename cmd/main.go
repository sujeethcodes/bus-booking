package main

import (
	"bus-booking/connectors"
	"os"

	"github.com/labstack/echo/v4"
)

func init() {
	connectors.LoadEnv()
}

func main() {
	e := echo.New()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	e.Start(":" + PORT)
}
