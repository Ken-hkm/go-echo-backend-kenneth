package api

import (
	"net/http"

	"github.com/Ken-hkm/go-echo-backend-kenneth/internal/db"
	"github.com/Ken-hkm/go-echo-backend-kenneth/internal/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

//func PersonalInfoHandler(c echo.Context) error {
//	collection := db.GetCollection("personal-info")
//	var info models.PersonalInfo
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	// Query the first document in the collection.
//	err := collection.FindOne(ctx, bson.M{}).Decode(&info)
//	if err != nil {
//		return c.JSON(http.StatusNotFound, map[string]string{"error": "Personal info not found"})
//	}
//
//	return c.JSON(http.StatusOK, info)
//}

func PersonalInfoHandler(c echo.Context) error {
	var info models.PersonalInfo
	filter := bson.M{"id": "67bdc56776af0fdb75bebfba"}

	documents, err := db.GetDocuments("personal-info", filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Check if at least one document was found
	if len(documents) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Personal info not found"})
	}

	// Convert BSON document into `models.PersonalInfo`
	bsonBytes, _ := bson.Marshal(documents[0])
	if err := bson.Unmarshal(bsonBytes, &info); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to decode document"})
	}

	// Return JSON response
	return c.JSON(http.StatusOK, info)
}
