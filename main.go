package main

import (
	"book-with-gin/controllers"
	"book-with-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})

	books := r.Group("/books")
	{
		books.GET("/", controllers.FindBooks)
		books.GET("/:id", controllers.FindBook)
		books.POST("/", controllers.CreateBook)
		books.PATCH("/:id", controllers.UpdateBook)
		books.DELETE("/:id", controllers.RemoveBook)
	}

	r.Run()
}
