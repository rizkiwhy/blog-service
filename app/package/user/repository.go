package user

import (
	"rizkiwhy-blog-service/package/user/model"
	"rizkiwhy-blog-service/util/database"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

type Repository interface {
	IsExistsByEmail(email string) (bool, error)
	Create(user model.User) (model.User, error)
}

type RepositoryImpl struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) Repository {
	return &RepositoryImpl{DB: DB}
}

func (r *RepositoryImpl) countByFilter(filter database.MySQLFilter) (res int64, err error) {
	db := database.BuildMySQLFilter(r.DB, filter)
	err = db.Model(&model.User{}).Count(&res).Error
	if err != nil {
		log.
			Err(err).
			Interface("filter", filter).
			Msg("[CountByFilter] Failed to count user by filter")
	}

	return
}

func (r *RepositoryImpl) IsExistsByEmail(email string) (res bool, err error) {
	filter := database.MySQLFilter{Where: gin.H{"email": email}}
	totalUsers, err := r.countByFilter(filter)
	if err != nil {
		log.Error().Err(err).Str("email", email).Msg("[IsExistsByEmail] Failed to count user by email")
		return
	}

	return totalUsers > 0, nil
}

func (r *RepositoryImpl) Create(user model.User) (model.User, error) {
	result := r.DB.Create(&user)
	if result.Error != nil {
		log.Error().Err(result.Error).Interface("user", user).Msg("[Create] Failed to create user")
	}

	return user, nil
}
