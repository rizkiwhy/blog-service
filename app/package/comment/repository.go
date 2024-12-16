package comment

import (
	"rizkiwhy-blog-service/package/comment/model"
	"rizkiwhy-blog-service/util/database"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Repository interface {
	Create(post model.Comment) (*model.Comment, error)
	SearchByFilter(filter database.MySQLFilter) (posts []model.Comment, err error)
}

type RepositoryImpl struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{DB: db}
}

func (r *RepositoryImpl) SearchByFilter(filter database.MySQLFilter) (comments []model.Comment, err error) {
	err = database.BuildMySQLFilter(r.DB, filter).Find(&comments).Error
	if err != nil {
		log.Error().Err(err).Interface("filter", filter).Msg("[CommentRepository][SearchByFilter] Failed to get comment by filter")
	}

	return
}

func (r *RepositoryImpl) Create(comment model.Comment) (*model.Comment, error) {
	result := r.DB.Create(&comment)
	if result.Error != nil {
		log.Error().Err(result.Error).Interface("comment", comment).Msg("[CommentRepository][Create] Failed to create comment")
		return nil, result.Error
	}

	return &comment, nil
}
