package api

import (
	"context"
	"net/http"
	"time"

	"github.com/Ken-hkm/go-echo-backend-kenneth/internal/db"
	"github.com/Ken-hkm/go-echo-backend-kenneth/internal/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func PersonalInfoHandler(c echo.Context) error {
	collection := db.GetCollection("personal-info")
	var info models.PersonalInfo

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Query the first document in the collection.
	err := collection.FindOne(ctx, bson.M{}).Decode(&info)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Personal info not found"})
	}

	return c.JSON(http.StatusOK, info)
}
