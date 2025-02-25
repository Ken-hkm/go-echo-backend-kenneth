package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CORSMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Set CORS headers on every response
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-KEY")

		// For preflight requests, return OK directly.
		if c.Request().Method == http.MethodOptions {
			return c.NoContent(http.StatusOK)
		}
		return next(c)
	}
}
