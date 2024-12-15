package post

import (
	"rizkiwhy-blog-service/package/post/model"
	"rizkiwhy-blog-service/util/database"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Repository interface {
	Create(post model.Post) (*model.Post, error)
	GetByID(id int64, preloadAuthor bool) (*model.Post, error)
	SearchByFilter(filter database.MySQLFilter) (posts []model.Post, err error)
	Update(post model.Post) (*model.Post, error)
	Delete(id int64) (err error)
}

type RepositoryImpl struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{DB: db}
}

func (r *RepositoryImpl) Delete(id int64) (err error) {
	result := r.DB.Delete(&model.Post{ID: id})
	if result.Error != nil {
		log.Error().Err(result.Error).Int64("id", id).Msg("[PostRepository][Delete] Failed to delete post")
		return result.Error
	}

	return nil
}

func (r *RepositoryImpl) Update(post model.Post) (*model.Post, error) {
	result := r.DB.Save(&post)
	if result.Error != nil {
		log.Error().Err(result.Error).Interface("post", post).Msg("[PostRepository][Update] Failed to update post")
		return nil, result.Error
	}

	return &post, nil
}

func (r *RepositoryImpl) Create(post model.Post) (*model.Post, error) {
	result := r.DB.Create(&post)
	if result.Error != nil {
		log.Error().Err(result.Error).Interface("post", post).Msg("[PostRepository][Create] Failed to create post")
		return nil, result.Error
	}

	return &post, nil
}

func (r *RepositoryImpl) GetByID(id int64, preloadAuthor bool) (post *model.Post, err error) {
	filter := database.MySQLFilter{Where: gin.H{"id": id}}
	if preloadAuthor {
		filter.Preload = []string{"Author"}
	}
	post, err = r.getByFilter(filter)
	if err != nil {
		log.Error().Err(err).Int64("id", id).Msg("[PostRepository][GetByID] Failed to get post by id")
	}

	return
}

func (r *RepositoryImpl) getByFilter(filter database.MySQLFilter) (post *model.Post, err error) {
	err = database.BuildMySQLFilter(r.DB, filter).First(&post).Error
	if err != nil {
		log.Error().Err(err).Interface("filter", filter).Msg("[PostRepository][getByFilter] Failed to get post by filter")
	}

	return
}

func (r *RepositoryImpl) SearchByFilter(filter database.MySQLFilter) (posts []model.Post, err error) {
	err = database.BuildMySQLFilter(r.DB, filter).Find(&posts).Error
	if err != nil {
		log.Error().Err(err).Interface("filter", filter).Msg("[PostRepository][SearchByFilter] Failed to get post by filter")
	}

	return
}
