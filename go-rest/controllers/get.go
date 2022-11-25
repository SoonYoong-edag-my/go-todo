package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rest/model"
	"log"
	"net/http"
)

func Get(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/todos", func(c *gin.Context) {
		log.Println("inside db var: ", db)
		var todos []model.Todo
		db.Find(&todos)
		c.IndentedJSON(http.StatusOK, todos)
	})
}
