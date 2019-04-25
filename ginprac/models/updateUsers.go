package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

//UpdateUsers ... Updates user details with new ones.
func UpdateUsers(c *gin.Context) {
	rawdata, _ := c.GetRawData()
	m := make(map[string]string)
	err := json.Unmarshal(rawdata, &m)
	if err != nil {
		fmt.Println(err)
	}
	for _, element := range users {
		if element.Name == c.Param("username") {
			age, _ := strconv.Atoi(m["age"])
			element = UserDetails{m["name"], age, m["email"], m["gender"], m["password"]}
		}
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Successfully updated user details for %s", c.Param("username")),
	})
}
