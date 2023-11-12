package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func HandleAuth(context echo.Context) error {
	err := context.Request().ParseForm()

	if err != nil {
		return err
	}

	username := context.Request().FormValue("username")
	password := context.Request().FormValue("password")

	// TODO: check username and password against database
	if username != viper.GetString("username") || password != viper.GetString("password") {
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
		"message": "Login successful",
		"token":   tokenString,
	})
}
