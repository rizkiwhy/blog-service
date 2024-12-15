package model

import (
	"rizkiwhy-blog-service/util/database"

	"github.com/gin-gonic/gin"
)

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

type Filter struct {
	Search string `json:"search"`
	Page   int64  `json:"page"`
	Limit  int64  `json:"limit"`
	Sort   string `json:"sort"`
	Order  string `json:"order"`
}

func (f *Filter) SetSearch(request string) {
	if request != "" {
		f.Search = request
	}
}

func (f *Filter) SetPagination(Page, Limit int64) {
	f.Page = Page
	f.Limit = Limit
}

func (f *Filter) SetSortAndOrder(Sort, Order string) {
	f.Sort = Sort
	f.Order = Order
}

func (f *Filter) ToMySQLFilter() database.MySQLFilter {
	return database.MySQLFilter{
		Like:   gin.H{"title": f.Search},
		Limit:  f.Limit,
		Offset: (f.Page - 1) * f.Limit,
		Order:  f.Sort,
		Sort:   f.Order,
	}
}
