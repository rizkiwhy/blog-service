package model

import (
	mPost "rizkiwhy-blog-service/package/post/model"
	mUser "rizkiwhy-blog-service/package/user/model"
	"time"
)

type Comment struct {
	ID        int64       `gorm:"primaryKey;autoIncrement"`
	PostID    int64       `gorm:"not null;index"`
	Post      *mPost.Post `gorm:"foreignKey:PostID"`
	AuthorID  int64       `gorm:"not null"`
	Author    *mUser.User `gorm:"foreignKey:AuthorID"`
	Content   string      `gorm:"type:text;not null"`
	CreatedAt time.Time   `gorm:"autoCreateTime;not null"`
}

func (c *Comment) ToCommentResponse() (res CommentResponse) {
	res = CommentResponse{
		ID:        c.ID,
		PostID:    c.PostID,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}

	if c.Author != nil {
		res.Author = &mUser.Author{
			ID:   c.Author.ID,
			Name: c.Author.Name,
		}
	}

	return
}
