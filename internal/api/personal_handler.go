package api

import (
	"github.com/Ken-hkm/go-echo-backend-kenneth/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func PersonalInfoHandler(c echo.Context) error {
	info := models.PersonalInfo{
		FirstName:   "Kenneth",
		LastName:    "Hakim",
		Email:       "hakimkenneth@gmail.com",
		Phone:       "+62-8788-8517-728",
		Address:     "Jl. Sunter Jaya VI B Blok L No 10, Jakarta Utara",
		LinkedInURL: "https://www.linkedin.com/in/kenneth-hakim-652b9612b/",
		GitHubURL:   "https://github.com/Ken-hkm",
	}
	// Add CORS headers
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Return the sample information with an HTTP 200 OK status.
	return c.JSON(http.StatusOK, info)
}
