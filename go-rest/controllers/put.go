package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rest/model"
	"net/http"
)

func Put(router *gin.Engine) {
	router.PUT("/todos/:id", func(c *gin.Context) {
		var existingTodo, updateTodo model.Todo

		if err := db.First(&existingTodo, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		if err := c.ShouldBindJSON(&updateTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updateTodo.Id = existingTodo.Id
		// have to use below instead of db.Model(&existingTodo).Updates(updateTodo) as it doesnt update to false or null value
		db.Model(&existingTodo).Updates(map[string]interface{}{"Name": updateTodo.Name, "Completed": updateTodo.Completed})
		c.IndentedJSON(http.StatusOK, existingTodo)
	})
}
