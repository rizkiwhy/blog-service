package post

import (
	"rizkiwhy-blog-service/package/post/model"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user model.Post) (*model.Post, error)
	GetByID(id int64) (*model.Post, error)
}

type RepositoryImpl struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{DB: db}
}

func (r *RepositoryImpl) Create(post model.Post) (*model.Post, error) {
	result := r.DB.Create(&post)
	if result.Error != nil {
		log.Error().Err(result.Error).Interface("post", post).Msg("[PostRepository][Create] Failed to create post")
		return nil, result.Error
	}

	return &post, nil
}

func (r *RepositoryImpl) GetByID(id int64) (post *model.Post, err error) {
	result := r.DB.Preload("Author").First(&post, id)
	if result.Error != nil {
		log.Error().Err(result.Error).Int64("id", id).Msg("[PostRepository][GetByID] Failed to get post")
		err = result.Error
	}

	return
}
