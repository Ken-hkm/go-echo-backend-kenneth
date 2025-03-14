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
	// Retrieve the "experience" collection from the database
	collection := db.GetCollection("experience")
	var experiences []models.Experience

	// Create a context with a 5-second timeout to prevent the database query from hanging indefinitely
	// The `defer cancel()` ensures resources are released even if the operation finishes early or encounters an error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response := utils.Response[interface{}]{
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve experiences",
			Data:       nil,
			Timestamp:  time.Now(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	defer cursor.Close(ctx)

	// Iterate through the cursor and decode each document
	for cursor.Next(ctx) {
		var exp models.Experience
		if err := cursor.Decode(&exp); err != nil {
			response := utils.Response[interface{}]{
				Status:     "error",
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to decode experiences",
				Data:       nil,
				Timestamp:  time.Now(),
			}
			return c.JSON(http.StatusInternalServerError, response)
		}
		experiences = append(experiences, exp)
	}

	if len(experiences) == 0 {
		response := utils.Response[interface{}]{
			Status:     "error",
			StatusCode: http.StatusNotFound,
			Message:    "No experiences found",
			Data:       nil,
			Timestamp:  time.Now(),
		}
		return c.JSON(http.StatusNotFound, response)
	}

	response := utils.Response[[]models.Experience]{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Experiences retrieved successfully",
		Data:       experiences,
		Timestamp:  time.Now(),
	}
	return c.JSON(http.StatusOK, response)
}
