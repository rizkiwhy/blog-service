package comment

import (
	"rizkiwhy-blog-service/package/comment/model"

	"github.com/rs/zerolog/log"
)

type ServiceImpl struct {
	Repository Repository
}

type Service interface {
	Create(request model.CreateRequest) (response model.CommentResponse, err error)
}

func NewService(repository Repository) Service {
	return &ServiceImpl{
		Repository: repository,
	}
}

func (s *ServiceImpl) Create(request model.CreateRequest) (response model.CommentResponse, err error) {
	comment := request.ToComment()
	result, err := s.Repository.Create(comment)
	if err != nil {
		log.Error().Err(err).Interface("comment", comment).Msg("[CommentService][Create] Failed to create comment")
		return
	}

	return result.ToCommentResponse(), nil
}
