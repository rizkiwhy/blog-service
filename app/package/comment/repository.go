package comment

import (
	"rizkiwhy-blog-service/package/comment/model"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Repository interface {
	Create(post model.Comment) (*model.Comment, error)
}

type RepositoryImpl struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{DB: db}
}

func (r *RepositoryImpl) Create(comment model.Comment) (*model.Comment, error) {
	result := r.DB.Create(&comment)
	if result.Error != nil {
		log.Error().Err(result.Error).Interface("comment", comment).Msg("[CommentRepository][Create] Failed to create comment")
		return nil, result.Error
	}

	return &comment, nil
}
