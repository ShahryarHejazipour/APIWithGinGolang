package controllers

import (
	"firstAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {

	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

//Create Book Validation
type CreateBookValidation struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {

	// Validate input
	var input CreateBookValidation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}

	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) {

	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

//Update Book Validation
type UpdateBookValidation struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {

	// Get model if exist
	var book models.Book

	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!!"})
		return
	}

	// Validate input
	var input UpdateBookValidation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
