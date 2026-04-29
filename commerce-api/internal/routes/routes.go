package routes

import (
	"commerce-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Register your routes here
	r.GET("/users", handlers.GetUser)
	r.GET("/products", handlers.GetProduct)
}
