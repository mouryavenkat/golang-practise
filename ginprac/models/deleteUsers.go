package models

import (
	"context"
	"fmt"
	"golang-practise/ginprac/mongoconn"
	"time"

	"github.com/gin-gonic/gin"
)

type dictionary map[string]string

//DeleteUsers ... Deletes users from DB based on userID, if not provided entirely.
func DeleteUsers(c *gin.Context) {
	client := mongoconn.Client
	collection := client.Database("demo").Collection("users")
	ctx, err := context.WithTimeout(context.Background(), 5*time.Second)
	if err != nil {
	}
	username := c.Query("username")
	var dbquery dictionary
	if username != "" {
		dbquery = dictionary{"name": username}
	} else {
		dbquery = dictionary{}
	}
	result, err2 := collection.DeleteMany(ctx, dbquery)
	if err2 != nil {
		fmt.Println("Error deleting items")
	}
	recordsDeleted := result.DeletedCount
	if recordsDeleted == 0 {
		c.JSON(200, gin.H{
			"message": "No records to delete",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Successfully deleted users",
	})
}
