package main

import (
	"github.com/Ken-hkm/go-echo-backend-kenneth/internal/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Define routes
	e.GET("/", handlers.HomeHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8085"))
}
