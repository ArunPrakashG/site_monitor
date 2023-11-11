package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func HandleAuth(context echo.Context) error {
	// get username and password from header
	// TODO: Get details from body
	username := context.Request().Header.Get("username")
	password := context.Request().Header.Get("password")

	// TODO: check username and password against database
	if username != "" || password != "" {
		return context.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid username or password",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	tokenString, err := token.SignedString([]byte("site-monitor"))

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
