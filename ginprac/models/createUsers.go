package models

import (
	"context"
	"fmt"
	"golang-practise/ginprac/mongoconn"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

//UserDetails ... describe the structure of user while registering
type UserDetails struct {
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// json:name is to make sure the incoming json is in bind with the struct

var users = []UserDetails{}

//CreateUser ... Used to create a new User
func CreateUser(c *gin.Context) {
	// Gets the data in the form of buffer, unmarshal it using json. The unmarshalled json will be stored in m.
	var userdetail UserDetails
	fmt.Println(c)
	binderr := c.ShouldBindJSON(&userdetail)
	fmt.Println(binderr)
	if binderr != nil {
		fmt.Println(binderr)
		c.JSON(500, gin.H{
			"message": "Input payload not matching",
			"error":   binderr,
		})
		return
	}
	fmt.Println(userdetail)

	Client := mongoconn.Client
	collection := Client.Database("demo").Collection("users")
	ctx, err1 := context.WithTimeout(context.Background(), 10*time.Second)
	if err1 != nil {
	}
	response, err2 := collection.InsertOne(ctx, userdetail)
	if err2 != nil {
		log.Println("Some error inserting the document")
	}
	fmt.Println(response.InsertedID)
	c.JSON(200, gin.H{
		"message": "User created successfully",
	})
}

// Other way of unmarshalling the response body

// buffer, _ := c.GetRawData()
// m := make(map[string]string)
// err := json.Unmarshal(buffer, &m)
// if err != nil {
// 	fmt.Println(err)
// 	c.JSON(500, gin.H{
// 		"message": "Unable to parse input payload",
// 		"error":   err,
// 	})
// 	return
// }
