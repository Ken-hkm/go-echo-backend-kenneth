package api

import (
	"github.com/Ken-hkm/go-echo-backend-kenneth/internal/middleware"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes sets up API routes.
func RegisterRoutes(e *echo.Echo) {
	// REST API Version 1 group
	v1 := e.Group("/api/v1")
	v1.Use(middleware.SecretKeyAuth)
	v1.GET("/personal-info", PersonalInfoHandler)
}
