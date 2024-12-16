package model

import (
	"time"

	mUser "rizkiwhy-blog-service/package/user/model"
)

const (
	ErrUnauthorizedAccess = "unauthorized access"
	ErrInvalidRequest     = "invalid request"
	ErrNotFound           = "record not found"
)

type PostResponse struct {
	ID        int64         `json:"id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	Author    *mUser.Author `json:"author,omitempty"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at,omitempty"`
}
