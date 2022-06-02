package main

import (
	"go-news/config"
	"go-news/controller"
	"go-news/repository"
	"go-news/service"

	"github.com/gin-gonic/gin"
)

func main() {

	DB := config.ConnectDB()

	newsRepository := repository.NewNewsRepository(DB)
	newsService := service.NewNewsService(newsRepository)
	newsController := controller.NewNewsController(newsService)

	router := gin.Default()

	router.GET("/api/news", newsController.FindAll)
	router.POST("/api/news", newsController.Create)
	router.GET("/api/news/:newsId", newsController.FindById)
	router.PUT("/api/news/:newsId", newsController.Update)
	router.DELETE("/api/news/:newsId", newsController.Delete)
	
	router.Run("localhost:8080")

}