package main

import (
	"golang-practise/ginprac/mongoconn"

	"github.com/gin-gonic/gin"
)

//Client ... Export mongo Client connection to all the packages

// c object for gin.Context is like response.json in node.
// Contains all the details about the request that the handler need to handle
func main() {
	// gin.Default logs every request and time taken to process the request ss well.
	//gin.New only does routing and rendering response by default. Any additional middleware has to be custom added.

	//router := gin.Default()
	mongoconn.ConnectToMongo()
	router := gin.New()
	// When used , will log every request details (time taken, request type etc)
	router.Use(gin.Logger())
	// When used gin.Recovery it catches an panic thrown by the code and by default return 500 statuscode.
	router.Use(gin.Recovery())

	routeDefinitions(router)
	router.Run(":8000")
}
