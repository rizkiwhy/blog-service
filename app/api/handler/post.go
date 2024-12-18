package handler

import (
	"net/http"
	"rizkiwhy-blog-service/api/presenter"
	pkgPost "rizkiwhy-blog-service/package/post"
	"rizkiwhy-blog-service/package/post/model"
	"rizkiwhy-blog-service/util/convert"
	"rizkiwhy-blog-service/util/database"

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
		c.JSON(http.StatusBadRequest, presenter.FailureResponse(presenter.CreatePostInvalidRequestMessage, err.Error()))
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

func (h *PostHandler) Update(c *gin.Context) {
	var request model.UpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("[PostHandler][Update] Failed to bind json")
		c.JSON(http.StatusBadRequest, presenter.FailureResponse(presenter.UpdatePostFailureMessage, err.Error()))
		return
	}
	request.ID = convert.StringToInt64(c.Param("id"))

	value, exists := c.Get("user_id")
	if !exists || value.(int64) == 0 {
		log.Error().Msg("[PostHandler][Update] Failed to get user id")
		c.JSON(http.StatusBadRequest, presenter.FailureResponse(presenter.UpdatePostFailureMessage, presenter.UpdatePostUnauthorizedMessage))
		return
	}
	request.AuthorID = value.(int64)

	response, err := h.Service.Update(request)
	if err != nil {
		log.Error().Err(err).Msg("[PostHandler][Update] Failed to update post")
		presenter.HandleError(c, err, presenter.UpdatePostStatusCodeMap, presenter.UpdatePostFailureMessage)
		return
	}

	c.JSON(http.StatusOK, presenter.SuccessResponse(presenter.UpdatePostSuccessMessage, response))
}

func (h *PostHandler) GetByID(c *gin.Context) {
	postID := convert.StringToInt64(c.Param("id"))
	response, err := h.Service.GetByID(postID)
	if err != nil {
		log.Error().Err(err).Msg("[PostHandler][Get] Failed to get post")
		presenter.HandleError(c, err, presenter.GetPostStatusCodeMap, presenter.GetPostFailureMessage)
		return
	}

	c.JSON(http.StatusOK, presenter.SuccessResponse(presenter.GetPostSuccessMessage, response))
}

func (h *PostHandler) Search(c *gin.Context) {
	var request database.Filter
	request.SetSearch(map[string][]string{c.Query("search"): {"content"}})
	request.SetPagination(convert.StringToInt64(c.DefaultQuery("page", "1")), convert.StringToInt64(c.DefaultQuery("limit", "10")))
	request.SetSortAndOrder(c.DefaultQuery("sort", "created_at"), c.DefaultQuery("order", "desc"))

	response, err := h.Service.SearchByFilter(request)
	if err != nil {
		log.Error().Err(err).Msg("[PostHandler][GetAll] Failed to get all post")
		presenter.HandleError(c, err, presenter.GetPostStatusCodeMap, presenter.GetPostFailureMessage)
		return
	}

	c.JSON(http.StatusOK, presenter.SuccessResponse(presenter.GetPostSuccessMessage, response))
}

func (h *PostHandler) Delete(c *gin.Context) {
	var request model.DeleteRequest
	request.ID = convert.StringToInt64(c.Param("id"))
	value, exists := c.Get("user_id")
	if !exists || value.(int64) == 0 {
		log.Error().Msg("[PostHandler][Update] Failed to get user id")
		c.JSON(http.StatusBadRequest, presenter.FailureResponse(presenter.UpdatePostFailureMessage, model.ErrUnauthorizedAccess))
		return
	}
	request.AuthorID = value.(int64)
	err := h.Service.Delete(request)
	if err != nil {
		log.Error().Err(err).Msg("[PostHandler][Delete] Failed to delete post")
		presenter.HandleError(c, err, presenter.DeletePostStatusCodeMap, presenter.DeletePostFailureMessage)
		return
	}

	c.JSON(http.StatusOK, presenter.SuccessResponse(presenter.DeletePostSuccessMessage, nil))
}
