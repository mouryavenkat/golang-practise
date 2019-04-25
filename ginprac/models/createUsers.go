package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

//UserDetails ... describe the structure of user while registering
type UserDetails struct {
	Name     string
	Age      int
	Email    string
	Gender   string
	Password string
}

var users = []UserDetails{}

//CreateUser ... Used to create a new User
func CreateUser(c *gin.Context) {
	// Gets the data in the form of buffer, unmarshal it using json. The unmarshalled json will be stored in m.
	buffer, _ := c.GetRawData()
	m := make(map[string]string)
	err := json.Unmarshal(buffer, &m)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": "Unable to parse input payload",
			"error":   err,
		})
		return
	}
	age, _ := strconv.Atoi(m["age"])
	userdetail := UserDetails{m["name"], age, m["email"], m["gender"], m["password"]}
	users = append(users, userdetail)
	for _, element := range users {
		fmt.Println(element)
	}
	c.JSON(200, gin.H{
		"message": "User created successfully",
	})
}
