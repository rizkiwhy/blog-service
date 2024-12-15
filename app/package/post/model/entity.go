package model

import (
	mUser "rizkiwhy-blog-service/package/user/model"
	"time"
)

type Post struct {
	ID        int64       `gorm:"primaryKey"`
	Title     string      `gorm:"type:varchar(255);not null"`
	Content   string      `gorm:"type:text;not null"`
	AuthorID  int64       `gorm:"not null"`
	Author    *mUser.User `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time   `gorm:"autoCreateTime;not null"`
	UpdatedAt *time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time  `gorm:"autoDeleteTime"`
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
		res.Author = &Author{
			ID:   p.Author.ID,
			Name: p.Author.Name,
		}
	}

	return
}
