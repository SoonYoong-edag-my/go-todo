package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rest/model"
	"net/http"
)

func Post(router *gin.Engine) {
	router.POST("/todos", func(c *gin.Context) {
		var newTodo model.Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&newTodo)
		c.IndentedJSON(http.StatusCreated, newTodo)
	})
}
