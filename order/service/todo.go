package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order/models"
)

func Create(c *gin.Context, todo *models.Todo) {
	err := models.CreateATodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
