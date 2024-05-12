package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID        string  `json:"id"`
	Writer    string  `json:"writer"`
	Title     string  `json:"title"`
	Publisher string  `json:"publisher"`
	Price     float64 `json:"price"`
}

var books = []book{
	{
		ID:        "1",
		Title:     "Clean Architecture: A Craftsman's Guide to Software Structure and Design",
		Writer:    "Robert C. Martin",
		Publisher: "Pearson; First Edition",
		Price:     120.99},
	{
		ID:        "2",
		Title:     "Refactoring: Improving the Design of Existing Code",
		Writer:    "Martin Fowler",
		Publisher: "Addison-Wesley Professional; Second Edition",
		Price:     49.99},
	{
		ID:        "3",
		Title:     "Cracking the Coding Interview",
		Writer:    "Gayle Laakmann McDowell",
		Publisher: "CareerCup; Sixth Edition",
		Price:     89.99},
	{
		ID:        "4",
		Title:     "The Art of Computer Programming",
		Writer:    "Donald Knuth and Donald John Fuller",
		Publisher: "Addison-Wesley Professional; First Edition",
		Price:     59.99},
	{
		ID:        "5",
		Title:     "A Philosophy of Software Design",
		Writer:    "John K. Ousterhout",
		Publisher: "Yaknyam Press",
		Price:     149.99},
	{
		ID:        "6",
		Title:     "The Pragmatic Programmer: Your Journey To Mastery, 20th Anniversary Edition",
		Writer:    "David Thomas, Andrew Hunt",
		Publisher: "Addison-Wesley Professional; Second Edition",
		Price:     189.99},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func postBooks(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)
	router.Run("localhost:8080")
}
