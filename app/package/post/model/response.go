package model

import "time"

const (
	ErrUnauthorizedAccess = "unauthorized access"
	ErrInvalidRequest     = "invalid request"
)

type PostResponse struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Author    *Author    `json:"author,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
