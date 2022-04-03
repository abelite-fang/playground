package routes

import (
	boardController "goland-api/controllers/v1/board"
	shortController "goland-api/controllers/v1/short"

	"github.com/gin-gonic/gin"
)

func BoardRoutes(router *gin.Engine) {
	//All routes related to users comes here
	v1 := router.Group("/v1")
	board := v1.Group("/board")
	short := v1.Group("/short")

	board.POST("/", boardController.CreateBoard())
	board.GET("/:name", boardController.GetBoards())

	short.POST("/", shortController.CreateShortenUrl())
}
