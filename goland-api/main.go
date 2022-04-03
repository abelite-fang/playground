package main

import (
	"goland-api/configs"
	"goland-api/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func routerEngine() *gin.Engine {
	// set server mode
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	configs.ConnectDB()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routes.BoardRoutes(r) //add this
	return r
}

func main() {
	//addr := ":" + os.Getenv("PORT")
	// log.Fatal(gateway.ListenAndServe("8080", routerEngine()))
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", routerEngine()))
}
