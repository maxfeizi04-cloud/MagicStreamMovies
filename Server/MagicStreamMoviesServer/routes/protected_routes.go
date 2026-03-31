package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/maxfeizi04-cloud/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/maxfeizi04-cloud/MagicStreamMovies/Server/MagicStreamMoviesServer/middleware"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
	router.GET("/recommendedmovies", controller.GetRecommendedMovies())

}
