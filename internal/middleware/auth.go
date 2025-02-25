package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SecretKeyAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		providedKey := c.Request().Header.Get("X-API-KEY")
		//hardcode, to maintain zero cost (since we are using aws lambda on free tier)
		//i did implement secrets feature but yeah, nope
		expectedKey := "wRNbm38KGBO79fj"

		if providedKey == "" || providedKey != expectedKey {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		}
		return next(c)
	}
}
