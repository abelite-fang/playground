package controllers

import (
	"context"
	"goland-api/configs"
	"goland-api/models"
	"goland-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "dyingapple-dev", "user")
var urlCollection *mongo.Collection = configs.GetCollection(configs.DB, "dyingapple-dev", "url")

var validate = validator.New()

func CreateShortenUrl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var urlRequest models.UrlRequest
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := c.BindJSON(&urlRequest); err != nil {
			c.JSON(http.StatusBadRequest,
				responses.GeneralResponse{
					Status:  http.StatusBadRequest,
					Message: "400 Bad Request: " + err.Error(),
				})
			return
		}

		if validationErr := validate.Struct(&urlRequest); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				responses.GeneralResponse{
					Status:  http.StatusBadRequest,
					Message: "Validate Error:" + validationErr.Error(),
				})
			return
		}

		result, err := urlCollection.InsertOne(ctx, urlRequest)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.GeneralResponse{
					Status:  http.StatusInternalServerError,
					Message: "Internal Error: " + err.Error(),
				})
			return
		}

		if id, ok := result.InsertedID.(primitive.ObjectID); ok {
			c.JSON(
				http.StatusCreated,
				responses.GeneralResponse{
					Status:  http.StatusCreated,
					Message: id.String(),
				})
		} else {
			println("Whattttt")
		}
	}
}
