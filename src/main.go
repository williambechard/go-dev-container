package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The Alchemist", Author: "Paulo Coelho", Quantity: 10},
	{ID: "2", Title: "The Little Prince", Author: "Antoine de Saint-Exupéry", Quantity: 5},
	{ID: "3", Title: "The Da Vinci Code", Author: "Dan Brown", Quantity: 7},
}

func main() {
	// setup router
	router := gin.Default()

	// setup routes
	setupRoutes(router)

	// run server
	router.Run("localhost:8080")
}

func setupRoutes(r *gin.Engine) {
	r.GET("/books", getBooks)
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
