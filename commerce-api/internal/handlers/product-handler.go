package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	// Implementation for getting product by Name
	products := []string{"Laptop", "Phone"}
	c.JSON(200, gin.H{
		"data": products,
	})
}
