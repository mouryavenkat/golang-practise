package models

import (
	"context"
	"fmt"
	"golang-practise/ginprac/mongoconn"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

//GetUserDetails ... Returns back the user Details
func GetUserDetails(c *gin.Context) {
	var listOfUsers = []UserDetails{}
	client := mongoconn.Client
	collection := client.Database("demo").Collection("users")
	ctx, err := context.WithTimeout(context.Background(), 5*time.Second)
	if err != nil {
	}
	userID := c.Query("username")
	var dbquery map[string]string
	if userID == "" {
		dbquery = map[string]string{}
	} else {
		dbquery = map[string]string{"name": userID}
	}
	log.Println("Get user details of ", userID)
	details, _ := collection.Find(ctx, dbquery)
	// Iterate until all the docs are done processing. Then Decode the result
	for details.Next(ctx) {
		var result UserDetails
		err := details.Decode(&result)
		if err != nil {
			fmt.Println("Some error decoding the userDetails object")
		}
		fmt.Println(result)
		listOfUsers = append(listOfUsers, result)
	}
	c.JSON(200, listOfUsers)
}
