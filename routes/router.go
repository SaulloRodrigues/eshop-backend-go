package routes

import (
	"eshop-backend-go/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/products", controllers.GetProducts)
	r.GET("/users", controllers.GetUsers)
	return r
}
