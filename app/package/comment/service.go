package comment

import (
	"rizkiwhy-blog-service/package/comment/model"
	"rizkiwhy-blog-service/util/database"

	"github.com/rs/zerolog/log"
)

type ServiceImpl struct {
	Repository Repository
}

type Service interface {
	Create(request model.CreateRequest) (response model.CommentResponse, err error)
	SearchByFilter(request database.Filter) (response []model.CommentResponse, err error)
}

func NewService(repository Repository) Service {
	return &ServiceImpl{
		Repository: repository,
	}
}

func (s *ServiceImpl) SearchByFilter(request database.Filter) (response []model.CommentResponse, err error) {
	filter := request.ToMySQLFilter()
	comments, err := s.Repository.SearchByFilter(filter)
	if err != nil {
		log.Error().Err(err).Interface("filter", filter).Msg("[CommentService][SearchByFilter] Failed to get comment by filter")
		return
	}

	for _, comment := range comments {
		response = append(response, comment.ToCommentResponse())
	}

	return
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
