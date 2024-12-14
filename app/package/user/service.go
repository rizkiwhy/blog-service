package user

import (
	"errors"
	"rizkiwhy-blog-service/package/user/model"

	"github.com/rs/zerolog/log"
)

type ServiceImpl struct {
	Repository Repository
}

type Service interface {
	Register(request model.RegisterRequest) (response model.RegisterResponse, err error)
}

func NewService(repository Repository) Service {
	return &ServiceImpl{
		Repository: repository,
	}
}

func (s *ServiceImpl) Register(request model.RegisterRequest) (response model.RegisterResponse, err error) {
	isExists, err := s.Repository.IsExistsByEmail(request.Email)
	if err != nil {
		log.Error().Err(err).Str("email", request.Email).Msg("[Register] Failed to count user by email")
		return
	}

	if isExists {
		err = errors.New(model.ErrEmailAlreadyExists)
		log.Error().Err(err).Str("email", request.Email).Msg("[Register] Email already exists")
		return
	}

	err = request.HashPassword()
	if err != nil {
		log.Error().Err(err).Msg("[Register] Failed to hash password")
		return
	}

	user := request.ToUser()

	result, err := s.Repository.Create(user)
	if err != nil {
		log.Error().Err(err).Interface("user", request).Msg("[Register] Failed to create user")
	}

	return result.ToRegisterResponse(), nil
}
