package handler

import (
	"net/http"
	"rizkiwhy-blog-service/api/presenter"
	pkgPost "rizkiwhy-blog-service/package/post"
	"rizkiwhy-blog-service/package/post/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PostHandler struct {
	Service pkgPost.Service
}

func NewPostHandler(service pkgPost.Service) *PostHandler {
	return &PostHandler{
		Service: service,
	}
}

func (h *PostHandler) Create(c *gin.Context) {
	var request model.CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("[PostHandler][Create] Failed to bind json")
		c.JSON(http.StatusBadRequest, presenter.FailureResponse(presenter.CreatePostFailureMessage, err.Error()))
		return
	}

	value, exists := c.Get("user_id")
	if !exists || value.(int64) == 0 {
		log.Error().Msg("[PostHandler][Create] Failed to get user id")
		c.JSON(http.StatusBadRequest, presenter.FailureResponse(presenter.CreatePostFailureMessage, presenter.CreatePostUnauthorizedMessage))
		return
	}
	request.AuthorID = value.(int64)

	response, err := h.Service.Create(request)
	if err != nil {
		log.Error().Err(err).Msg("[PostHandler][Create] Failed to create post")
		presenter.HandleError(c, err, presenter.CreatePostStatusCodeMap, presenter.CreatePostFailureMessage)
		return
	}

	c.JSON(http.StatusCreated, presenter.SuccessResponse(presenter.CreatePostSuccessMessage, response))
}
