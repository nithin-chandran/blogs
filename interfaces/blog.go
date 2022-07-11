package interfaces

import "blogs/models"

// Blog inetrface for all the functions related to blog.
type Blog interface {
	Create(models.BlogModel) error
	GetAll(limit int, offset int) []models.BlogModel
	GetBlogs(id int) models.BlogModel
	DeleteBlog(id int)
	GetRecentBlogs()
}
