package controllers

import (
	"go_mongo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.InsertMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert movie"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Movie inserted successfully!"})
}

func UpdateMovie(c *gin.Context) {
	movieId := c.Param("id")
	var movie models.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.UpdateMovie(movieId, movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully!"})
}

func DeleteMovie(c *gin.Context) {
	movieId := c.Param("id")

	err := models.DeleteMovie(movieId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete movie"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully!"})
}

func DeleteAllMovies(c *gin.Context) {
	err := models.DeleteAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete all movies"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All movies deleted successfully!"})
}

func ListAllMovies(c *gin.Context) {
	movies := models.ListAll()
	c.JSON(http.StatusOK, movies)
}

func FindMovieByName(c *gin.Context) {
	movieName := c.Param("name")
	movie := models.Find(movieName)
	if movie.Movie == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func FindAllMoviesByName(c *gin.Context) {
	movieName := c.Param("name")
	movies := models.FindAll(movieName)
	if len(movies) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No movies found"})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func InsertMultipleMovies(c *gin.Context) {
	var movies []models.Movie
	if err := c.ShouldBindJSON(&movies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.InsertMany(movies)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert movies"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Movies inserted successfully!"})
}
