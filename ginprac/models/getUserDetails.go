package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//GetUserDetails ... Returns back the user Details
func GetUserDetails(c *gin.Context) {
	fmt.Println(users)

	c.JSON(200, users)
}
