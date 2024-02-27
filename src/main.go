package main

import (
	"errors"
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
	r.GET("/books/:id", bookById)
	r.POST("/books", createBook)
	r.PATCH("/checkout", checkoutBook)
	r.PATCH("/return", returnBook)
}

func bookById(c *gin.Context) {
	id := c.Param("id")

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id parameter"})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not available"})
		return
	}

	book.Quantity--
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id parameter"})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Quantity++
	c.IndentedJSON(http.StatusOK, book)

}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
