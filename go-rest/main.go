package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rest/lib"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	lib.Demo()
	router := gin.Default()
	router.Use(CORSMiddleware())

	// Binding from JSON
	type Todo struct {
		Id        uint   `json:"id" gorm:"primary_key"`
		Name      string `json:"name" binding:"required"`
		Completed bool   `json:"completed"`
	}

	// init db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	if err = db.AutoMigrate(&Todo{}); err != nil {
		panic("failed to create schema")
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/todos", func(c *gin.Context) {
		var todos []Todo
		db.Find(&todos)
		c.IndentedJSON(http.StatusOK, todos)
	})

	router.POST("/todos", func(c *gin.Context) {
		var newTodo Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&newTodo)
		c.IndentedJSON(http.StatusCreated, newTodo)
	})

	router.PUT("/todos/:id", func(c *gin.Context) {
		var existingTodo, updateTodo Todo

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

	router.DELETE("/todos/:id", func(c *gin.Context) {
		var existingTodo Todo
		if err := db.First(&existingTodo, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		db.Model(&existingTodo).Delete(&Todo{}, existingTodo.Id)
	})

	router.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
