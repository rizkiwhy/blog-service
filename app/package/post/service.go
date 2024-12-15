package post

import (
	"rizkiwhy-blog-service/package/post/model"

	"github.com/rs/zerolog/log"
)

type ServiceImpl struct {
	Repository Repository
}

type Service interface {
	Create(request model.CreateRequest) (response model.PostResponse, err error)
}

func NewService(repository Repository) Service {
	return &ServiceImpl{
		Repository: repository,
	}
}

func (s *ServiceImpl) Create(request model.CreateRequest) (response model.PostResponse, err error) {
	post := request.ToPost()
	result, err := s.Repository.Create(post)
	if err != nil {
		log.Error().Err(err).Interface("post", post).Msg("[PostService][Create] Failed to create post")
		return
	}

	return result.ToPostResponse(), nil
}
