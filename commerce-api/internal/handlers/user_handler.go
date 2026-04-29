package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	// Implementation for getting user by ID
	users := []string{"Alice", "Bob"}
	c.JSON(200, gin.H{
		"data": users,
	})
}
