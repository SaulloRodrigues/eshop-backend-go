package routes

import (
	"eshop-backend-go/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
	}
}