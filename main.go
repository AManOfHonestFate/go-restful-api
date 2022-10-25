package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type film struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Director string  `json:"director"`
	Price    float32 `json:"price"`
}

var films = []film{
	{ID: "1", Title: "The Hateful Eight", Director: "Quentin Tarantino", Price: 56.99},
	{ID: "2", Title: "Eraserhead", Director: "David Lynch", Price: 17.99},
	{ID: "3", Title: "No Country For Old Man", Director: "Coen brothers", Price: 39.99},
}

func getFilms(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, films)
}

func postFilms(c *gin.Context) {
	var newFilm film

	if err := c.BindJSON(&newFilm); err != nil {
		return
	}

	films = append(films, newFilm)
	c.IndentedJSON(http.StatusCreated, newFilm)
}

func main() {
	router := gin.Default()
	router.GET("/films", getFilms)
	router.POST("/films", postFilms)

	router.Run("localhost:8000")
}
