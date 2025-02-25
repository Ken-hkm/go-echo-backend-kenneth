package middleware

import (
	"github.com/Ken-hkm/go-echo-backend-kenneth/internal/secrets"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func SecretKeyAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		providedKey := c.Request().Header.Get("X-API-KEY")
		expectedKey, err := secrets.GetParameter("go_api_key")
		if err != nil {
			log.Printf("Failed to get parameter: %v", err)
			errMsg := "Failed to get parameter: " + err.Error()
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": errMsg})
		}
		if providedKey == "" || providedKey != expectedKey {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		}
		return next(c)
	}
}
