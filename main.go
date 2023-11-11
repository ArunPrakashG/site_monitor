package main

import (
	"site_monitor/src/auth"
	"site_monitor/src/endpoints"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/auth", auth.HandleAuth)
	e.GET("/stats", endpoints.HandleStats, echojwt.JWT([]byte("site-monitor")))
	e.Logger.Fatal(e.Start(":1111"))
}
