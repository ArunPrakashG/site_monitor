package endpoints

import (
	"net/http"
	monitor "site_monitor/src/monitors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UnauthorizedResponse struct {
	Message       string `json:"message"`
	ReceivedToken string `json:"token"`
}

type StatsResponse struct {
	CPU monitor.CPU `json:"cpu"`
	RAM monitor.RAM `json:"ram"`
	// Processes monitor.Processes `json:"processes"`
	Disks monitor.Disks `json:"disk"`
}

func HandleStats(context echo.Context) error {
	token, ok := context.Get("user").(*jwt.Token)

	if !ok {
		return context.JSONPretty(http.StatusUnauthorized, UnauthorizedResponse{
			Message:       "JWT token missing or invalid",
			ReceivedToken: token.Raw,
		}, "  ")
	}

	return context.JSON(http.StatusOK, StatsResponse{
		CPU:   monitor.CheckCPU(),
		RAM:   monitor.CheckRAM(),
		Disks: monitor.CheckDisks(),
	})

}
