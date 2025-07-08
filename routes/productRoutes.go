package routes 

import (
	"eshop-backend-go/controllers"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(rg *gin.RouterGroup) {
	productGroup := rg.Group("/products")
	{
		productGroup.GET("/", controllers.GetProducts)
	}
}