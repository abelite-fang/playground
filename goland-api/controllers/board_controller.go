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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var boardCollection *mongo.Collection = configs.GetCollection(configs.DB, "dyingapple-dev", "board")
var cardCollection *mongo.Collection = configs.GetCollection(configs.DB, "dyingapple-dev", "card")
var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "dyingapple-dev", "user")
var validate = validator.New()

func GetBoards() gin.HandlerFunc {
	return func(c *gin.Context) {
		var boards []*models.Board
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var query = c.Param("name")
		cursor, err := boardCollection.Find(
			context.Background(),
			bson.M{"name": query},
			&options.FindOptions{},
		)

		if err != nil {
			panic(err)
		}

		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			board := &models.Board{}
			if err := cursor.Decode(board); err != nil {
				panic(err)
			}
			boards = append(boards, board)
		}

		c.JSON(
			http.StatusOK,
			boards,
		)
	}
}
func CreateBoard() gin.HandlerFunc {
	return func(c *gin.Context) {
		var board models.Board
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := c.BindJSON(&board); err != nil {
			c.JSON(http.StatusBadRequest,
				responses.GeneralResponse{
					Status:  http.StatusBadRequest,
					Message: "400 Bad Request",
				})
			return
		}

		if validationErr := validate.Struct(&board); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				responses.GeneralResponse{
					Status:  http.StatusBadRequest,
					Message: "Validate Error:" + validationErr.Error(),
				})
			return
		}

		newBoard := models.Board{
			Id:   primitive.NewObjectID(),
			Name: board.Name,
		}

		_, err := boardCollection.InsertOne(ctx, newBoard)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.GeneralResponse{
					Status:  http.StatusInternalServerError,
					Message: "Internal Error: " + err.Error(),
				})
			return
		}

		c.JSON(
			http.StatusCreated,
			responses.GeneralResponse{
				Status:  http.StatusCreated,
				Message: "success",
			})
	}
}
