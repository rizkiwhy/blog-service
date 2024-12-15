package model

import (
	"errors"
	mUser "rizkiwhy-blog-service/package/user/model"
	"time"

	"github.com/rs/zerolog/log"
)

type Post struct {
	ID        int64       `gorm:"primaryKey"`
	Title     string      `gorm:"type:varchar(255);not null"`
	Content   string      `gorm:"type:text;not null"`
	AuthorID  int64       `gorm:"not null"`
	Author    *mUser.User `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time   `gorm:"autoCreateTime;not null"`
	UpdatedAt *time.Time  `gorm:"autoUpdateTime"`
}

func (p *Post) ValidateAuthor(userID int64) bool {
	return p.Author.ID == userID
}

func (p *Post) UpdateRequest(request UpdateRequest) (err error) {
	if request.AuthorID != p.AuthorID {
		err = errors.New(ErrUnauthorizedAccess)
		log.Error().Int64("author_id", request.AuthorID).Int64("post_author_id", p.AuthorID).Msg("[Post][UpdateRequest] Unauthorized")
		return
	}

	if request.Title != "" {
		if request.Title != p.Title {
			p.Title = request.Title
		}
	}

	if request.Content != "" {
		if request.Content != p.Content {
			p.Content = request.Content
		}
	}

	return
}

func (p *Post) ToPostResponse() (res PostResponse) {
	res = PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}

	if p.Author != nil {
		res.Author = &mUser.Author{
			ID:   p.Author.ID,
			Name: p.Author.Name,
		}
	}

	return
}
