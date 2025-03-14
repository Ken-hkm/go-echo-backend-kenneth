package handlers

import (
	"github.com/Ken-hkm/go-echo-backend-kenneth/internal/db"
	"github.com/Ken-hkm/go-echo-backend-kenneth/internal/models"
	"github.com/Ken-hkm/go-echo-backend-kenneth/internal/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

func ExperienceHandler(c echo.Context) error {
	collection := db.GetCollection("experience")
	var info models.Experience

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Query the first document in the collection.
	err := collection.FindOne(ctx, bson.M{}).Decode(&info)
	if err != nil {
		response := utils.Response[interface{}]{
			Status:     "error",
			StatusCode: http.StatusNotFound,
			Message:    "Experience not found",
			Data:       nil,
			Timestamp:  time.Now(),
		}
		return c.JSON(http.StatusNotFound, response)
	}

	response := utils.Response[models.Experience]{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Experience retrieved successfully",
		Data:       info,
		Timestamp:  time.Now(),
	}
	return c.JSON(http.StatusOK, response)
}
