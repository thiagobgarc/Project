// main.go
package main

//import of folders and packages
import (
	"commerce-api/internal/routes"

	"github.com/gin-gonic/gin"
)

// @title Commerce API
// @version 1.0
// @description This is a sample server for a commerce API.
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
