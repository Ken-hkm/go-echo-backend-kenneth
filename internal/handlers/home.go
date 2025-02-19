package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello from Echo on AWS Lambda!"})
}
