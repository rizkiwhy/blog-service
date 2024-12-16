package model

import (
	mUser "rizkiwhy-blog-service/package/user/model"
	"time"
)

const (
	ErrUnauthorizedAccess = "unauthorized access"
	ErrInvalidRequest     = "invalid request"
	ErrNotFound           = "record not found"
)

type CommentResponse struct {
	ID        int64         `json:"id"`
	PostID    int64         `json:"post_id"`
	Author    *mUser.Author `json:"author,omitempty"`
	Content   string        `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at,omitempty"`
}
