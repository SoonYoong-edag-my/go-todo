package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rest/model"
	"net/http"
)

func Delete(router *gin.Engine) {
	router.DELETE("/todos/:id", func(c *gin.Context) {
		var existingTodo model.Todo
		if err := db.First(&existingTodo, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		db.Model(&existingTodo).Delete(&model.Todo{}, existingTodo.Id)
	})
}
