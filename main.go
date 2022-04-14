package main

import (
	"firstAPI/controllers"
	"firstAPI/models"
	"github.com/gin-gonic/gin"
)

func main() {

	rout := gin.Default()

	models.ConnectDatabase()

	rout.GET("/books", controllers.FindBooks)
	rout.POST("/books", controllers.CreateBook)
	rout.GET("/books/:id", controllers.FindBook)
	rout.PATCH("/books/:id", controllers.UpdateBook)
	rout.DELETE("/books/:id", controllers.DeleteBook)
	rout.Run()

}
