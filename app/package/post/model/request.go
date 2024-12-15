package model

type CreateRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	AuthorID int64  `json:"author_id"`
}

func (r *CreateRequest) ToPost() Post {
	return Post{
		Title:    r.Title,
		Content:  r.Content,
		AuthorID: r.AuthorID,
	}
}
