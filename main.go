package main

import (
	"fmt"
	config "site_monitor/src"
	"site_monitor/src/auth"
	"site_monitor/src/endpoints"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	isSuccess, err := config.Load()

	if !isSuccess || err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/auth", auth.HandleAuth)
	e.GET("/stats", endpoints.HandleStats, echojwt.JWT([]byte(viper.GetString("secret"))))
	e.Logger.Fatal(e.Start(":" + fmt.Sprint(viper.GetInt("port"))))
}
