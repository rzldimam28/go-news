package controller

import (
	"go-news/helper"
	"go-news/service"
	"go-news/web/req"
	"go-news/web/res"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewsController interface {
	FindAll(c *gin.Context)
	Create(c *gin.Context)
	FindById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type NewsControllerImpl struct {
	NewsService service.NewsService
}

func NewNewsController(newsService service.NewsService) NewsController {
	return &NewsControllerImpl{
		NewsService: newsService,
	}
}

func (newsController *NewsControllerImpl) FindAll(c *gin.Context) {
	newsResponses := newsController.NewsService.FindAll(c)
	webResponse := res.WebResponse{
		Code: http.StatusOK,
		Status: "Success Get All News",
		Data: newsResponses,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (newsController *NewsControllerImpl) Create(c *gin.Context) {
	newsRequest := req.NewsCreateRequest{}
	if err := c.ShouldBindJSON(&newsRequest); err != nil {
		helper.ErrorCantProcessRequest(c)
		return
	}
	newsResponse := newsController.NewsService.Create(c, newsRequest)
	webResponse := res.WebResponse{
		Code: http.StatusCreated,
		Status: "Success Create New News",
		Data: newsResponse,
	}
	c.JSON(http.StatusCreated, webResponse)
}

func (newsController *NewsControllerImpl) FindById(c *gin.Context) {
	id := c.Param("newsId")
	newsId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helper.ErrorNotFound(c)
		return
	}
	newsResponse := newsController.NewsService.FindById(c, newsId)
	webResponse := res.WebResponse{
		Code: http.StatusOK,
		Status: "Success Get News By Id",
		Data: newsResponse,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (newsController *NewsControllerImpl) Update(c *gin.Context) {
	id := c.Param("newsId")
	newsId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helper.ErrorNotFound(c)
		return
	}

	newsRequest := req.NewsUpdateRequest{}
	if err := c.ShouldBindJSON(&newsRequest); err != nil {
		helper.ErrorCantProcessRequest(c)
		return
	}

	newsRequest.Id = newsId
	newsResponse := newsController.NewsService.Update(c, newsRequest)
	webResponse := res.WebResponse{
		Code: http.StatusOK,
		Status: "Success Update News By Id",
		Data: newsResponse,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (newsController *NewsControllerImpl) Delete(c *gin.Context) {
	id := c.Param("newsId")
	newsId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helper.ErrorNotFound(c)
		return
	}
	newsController.NewsService.Delete(c, newsId)
	webResponse := res.WebResponse{
		Code: 200,
		Status: "Success Delete News By Id",
		Data: nil,
	}
	c.JSON(http.StatusOK, webResponse)
}