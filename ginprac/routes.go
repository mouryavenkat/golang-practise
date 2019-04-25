package main

import (
	"fmt"
	"golang-practise/ginprac/models"

	"github.com/gin-gonic/gin"
)

func middleware(c *gin.Context) {
	fmt.Println("testing middleware")
}

func routeDefinitions(router *gin.Engine) {
	router.GET("/healthCheck", func(c *gin.Context) {
		middleware(c)
	}, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"health":  "Awesome",
			"version": "1.0",
		})
	})
	router.POST("/createUser", func(c *gin.Context) {
		models.CreateUser(c)
	})
	router.GET("/getUserDetails", func(c *gin.Context) {
		models.GetUserDetails(c)
	})
	router.PUT("/updateUser/:username", func(c *gin.Context) {
		models.UpdateUsers(c)
	})
}
