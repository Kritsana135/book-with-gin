package controllers

import (
	"book-with-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func CreateBook(c *gin.Context) {
	var bookInput CreateBookInput
	if err := c.ShouldBindJSON(&bookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: bookInput.Title, Author: bookInput.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var bookInput models.Book
	if err := c.ShouldBindJSON(&bookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	// models.DB.First(&book, id)
	if err := models.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
		return
	}

	models.DB.Model(&book).Update(bookInput)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func RemoveBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := models.DB.Delete(&book, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Book deleted successfully"})
}
