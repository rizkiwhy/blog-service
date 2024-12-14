package handler

import (
	"net/http"
	"rizkiwhy-blog-service/api/presenter"
	pkgUser "rizkiwhy-blog-service/package/user"
	"rizkiwhy-blog-service/package/user/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserHandler struct {
	Service pkgUser.Service
}

func NewUserHandler(service pkgUser.Service) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var request model.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("[UserHandler][Register] Failed to bind json")
		c.JSON(http.StatusBadRequest, presenter.FailureResponse(presenter.RegisterInvalidRequestMessage, err.Error()))
		return
	}

	response, err := h.Service.Register(request)
	if err != nil {
		log.Error().Err(err).Msg("[UserHandler][Register] Failed to register user")
		c.JSON(presenter.RegisterStatusCodeMap[err.Error()], presenter.FailureResponse(presenter.RegisterFailureMessage, err.Error()))
		return
	}

	c.JSON(http.StatusOK, presenter.SuccessResponse(presenter.RegisterSuccessMessage, response))
}

func (h *UserHandler) Login(c *gin.Context) {
	var request model.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("[UserHandler][Login] Failed to bind json")
		c.JSON(http.StatusBadRequest, presenter.FailureResponse(presenter.LoginInvalidCredentialsMessage, err.Error()))
		return
	}

	response, err := h.Service.Login(request)
	if err != nil {
		log.Error().Err(err).Msg("[UserHandler][Login] Failed to login user")
		c.JSON(presenter.LoginStatusCodeMap[err.Error()], presenter.FailureResponse(presenter.LoginFailureMessage, err.Error()))
		return
	}

	c.JSON(http.StatusOK, presenter.SuccessResponse(presenter.LoginSuccessMessage, response))
}
