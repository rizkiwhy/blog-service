package model

type DeleteRequest struct {
	ID       int64
	AuthorID int64
}

type UpdateRequest struct {
	ID       int64
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int64
}

type CreateRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	AuthorID int64
}

func (r *CreateRequest) ToPost() Post {
	return Post{
		Title:    r.Title,
		Content:  r.Content,
		AuthorID: r.AuthorID,
	}
}
