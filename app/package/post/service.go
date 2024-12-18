package post

import (
	"errors"
	"rizkiwhy-blog-service/package/post/model"
	"rizkiwhy-blog-service/util/database"

	"github.com/rs/zerolog/log"
)

type ServiceImpl struct {
	Repository Repository
}

type Service interface {
	Create(request model.CreateRequest) (response model.PostResponse, err error)
	GetByID(request int64) (response model.PostResponse, err error)
	SearchByFilter(request database.Filter) (response []model.PostResponse, err error)
	Update(request model.UpdateRequest) (response model.PostResponse, err error)
	Delete(request model.DeleteRequest) (err error)
}

func NewService(repository Repository) Service {
	return &ServiceImpl{
		Repository: repository,
	}
}

func (s *ServiceImpl) Delete(request model.DeleteRequest) (err error) {
	post, err := s.Repository.GetByID(request.ID, true)
	if err != nil {
		log.Error().Err(err).Int64("id", request.ID).Msg("[PostService][Delete] Failed to get post by id")
		return
	}

	if !post.ValidateAuthor(request.AuthorID) {
		err = errors.New(model.ErrUnauthorizedAccess)
		log.Error().Int64("author_id", request.AuthorID).Int64("post_author_id", post.AuthorID).Msg("[PostService][Delete] Unauthorized")
		return
	}

	err = s.Repository.Delete(request.ID)
	if err != nil {
		log.Error().Err(err).Int64("id", request.ID).Msg("[PostService][Delete] Failed to delete post")
	}

	return
}

func (s *ServiceImpl) Update(request model.UpdateRequest) (response model.PostResponse, err error) {
	post, err := s.Repository.GetByID(request.ID, false)
	if err != nil {
		log.Error().Err(err).Int64("id", request.ID).Msg("[PostService][Update] Failed to get post by id")
		return
	}

	err = post.UpdateRequest(request)
	if err != nil {
		log.Error().Err(err).Interface("post", post).Msg("[PostService][Update] Failed to update request")
		return
	}

	result, err := s.Repository.Update(*post)
	if err != nil {
		log.Error().Err(err).Interface("post", post).Msg("[PostService][Update] Failed to update post")
		return
	}

	return result.ToPostResponse(), nil
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

func (s *ServiceImpl) GetByID(request int64) (response model.PostResponse, err error) {
	post, err := s.Repository.GetByID(request, true)
	if err != nil {
		log.Error().Err(err).Int64("id", request).Msg("[PostService][GetByID] Failed to get post by id")
		return
	}

	return post.ToPostResponse(), nil
}

func (s *ServiceImpl) SearchByFilter(request database.Filter) (response []model.PostResponse, err error) {
	filter := request.ToMySQLFilter()
	filter.Preload = append(filter.Preload, "Author")
	posts, err := s.Repository.SearchByFilter(filter)
	if err != nil {
		log.Error().Err(err).Interface("filter", filter).Msg("[PostService][SearchByFilter] Failed to get post by filter")
		return
	}

	for _, post := range posts {
		response = append(response, post.ToPostResponse())
	}

	return
}
