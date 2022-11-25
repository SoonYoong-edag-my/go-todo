package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rest/controllers"
	"log"
)

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	controllers.Get(router)
	controllers.Post(router)
	controllers.Delete(router)
	controllers.Put(router)

	err := router.Run(":8080")
	if err != nil {
		log.Println("unable to run at port 8080")
	}
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
