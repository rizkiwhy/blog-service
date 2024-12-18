package handler

import (
	"net/http"
	"rizkiwhy-blog-service/api/presenter"
	pkgComment "rizkiwhy-blog-service/package/comment"
	"rizkiwhy-blog-service/package/comment/model"
	"rizkiwhy-blog-service/util/convert"
	"rizkiwhy-blog-service/util/database"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CommentHandler struct {
	Service pkgComment.Service
}

func NewCommentHandler(service pkgComment.Service) *CommentHandler {
	return &CommentHandler{
		Service: service,
	}
}

func (h *CommentHandler) Search(c *gin.Context) {
	var request database.Filter
	request.SetPagination(convert.StringToInt64(c.DefaultQuery("page", "1")), convert.StringToInt64(c.DefaultQuery("limit", "10")))
	request.SetSortAndOrder(c.DefaultQuery("sort", "created_at"), c.DefaultQuery("order", "desc"))
	postID := convert.StringToInt64(c.Param("id"))
	request.Equal = gin.H{"post_id": postID}
	response, err := h.Service.SearchByFilter(request)
	if err != nil {
		log.Error().Err(err).Msg("[CommentHandler][GetAll] Failed to get all comment")
		presenter.HandleError(c, err, presenter.GetCommentStatusCodeMap, presenter.GetCommentFailureMessage)
		return
	}

	c.JSON(http.StatusOK, presenter.SuccessResponse(presenter.GetCommentSuccessMessage, response))
}

func (h *CommentHandler) Create(c *gin.Context) {
	var request model.CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("[CommentHandler][Create] Failed to bind json")
		c.JSON(http.StatusBadRequest, presenter.FailureResponse(presenter.CreateCommentInvalidRequestMessage, err.Error()))
		return
	}

	value, exists := c.Get("user_id")
	if !exists || value.(int64) == 0 {
		log.Error().Msg("[CommentHandler][Create] Failed to get user id")
		c.JSON(http.StatusBadRequest, presenter.FailureResponse(presenter.CreateCommentFailureMessage, presenter.CreateCommentUnauthorizedMessage))
		return
	}

	request.PostID = convert.StringToInt64(c.Param("id"))
	request.AuthorID = value.(int64)
	if err := request.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, presenter.FailureResponse(presenter.CreateCommentInvalidRequestMessage, err.Error()))
		return
	}

	response, err := h.Service.Create(request)
	if err != nil {
		log.Error().Err(err).Msg("[CommentHandler][Create] Failed to create comment")
		presenter.HandleError(c, err, presenter.CreateCommentStatusCodeMap, presenter.CreateCommentFailureMessage)
		return
	}

	c.JSON(http.StatusCreated, presenter.SuccessResponse(presenter.CreateCommentSuccessMessage, response))
}
