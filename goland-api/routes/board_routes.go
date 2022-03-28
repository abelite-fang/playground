package routes

import (
	"goland-api/controllers"

	"github.com/gin-gonic/gin"
)

func BoardRoutes(router *gin.Engine) {
	//All routes related to users comes here
	router.POST("/board", controllers.CreateBoard())    //add this
	router.GET("/board/:name", controllers.GetBoards()) //add this

}
