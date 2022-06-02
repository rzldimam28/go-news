package helper

import (
	"go-news/web/res"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorNotFound(c *gin.Context) {
	webResponse := res.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "Can Not Found News ID",
		Data:   nil,
	}
	c.JSON(http.StatusBadRequest, webResponse)
}

func ErrorCantProcessRequest(c *gin.Context) {
	webResponse := res.WebResponse{
		Code: http.StatusBadRequest,
		Status: "Can Not Process From Request Body",
		Data: nil,
	}
	c.JSON(http.StatusBadRequest, webResponse)
}