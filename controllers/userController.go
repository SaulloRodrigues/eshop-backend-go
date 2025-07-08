package controllers 

import (
	"eshop-backend-go/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password123"},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Password: "password456"},
	}
	c.JSON(200, users)
}