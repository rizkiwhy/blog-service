package model

import "errors"

type CreateRequest struct {
	PostID   int64
	AuthorID int64
	Content  string `json:"content" binding:"required"`
}

func (r *CreateRequest) ToComment() Comment {
	return Comment{
		PostID:   r.PostID,
		AuthorID: r.AuthorID,
		Content:  r.Content,
	}
}

func (r *CreateRequest) Validate() (err error) {
	if r.PostID == 0 || r.AuthorID == 0 || r.Content == "" {
		err = errors.New(ErrInvalidRequest)
	}

	return
}
