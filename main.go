package main

import (
	"go_mongo/controllers"
	"go_mongo/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Init session store
	r.Use(gin.Logger())

	models.ConnectDatabase()

	movies := r.Group("/movies")
	{
		movies.POST("/", controllers.CreateMovie)
		movies.PUT("/:id", controllers.UpdateMovie)
		movies.DELETE("/:id", controllers.DeleteMovie)
		movies.DELETE("/", controllers.DeleteAllMovies)
		movies.GET("/", controllers.ListAllMovies)
		movies.GET("/one/:name", controllers.FindMovieByName)
		movies.GET("/all/:name", controllers.FindAllMoviesByName)
		movies.POST("/multiple", controllers.InsertMultipleMovies)
	}

	log.Println("Server started!")
	r.Run() // Default Port 8080
}
