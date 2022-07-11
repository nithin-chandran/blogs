package service

import (
	"blogs/interfaces"
	"blogs/models"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

// JSONBlogService implements the blog interface for json storage implementations.
type JSONBlogService struct {
	cache interfaces.Cache
}

var filePermissionMode fs.FileMode = 0644

// Create creates the blog.
func (blogService *JSONBlogService) Create(blogModel models.BlogModel) error {
	jsonData, err := os.ReadFile("data/store.json")
	if err != nil {
		return fmt.Errorf("error on raeding file")
	}
	existingBlogs := []models.BlogModel{}
	if len(jsonData) > 0 {
		err = json.Unmarshal(jsonData, &existingBlogs)
		if err != nil {
			return fmt.Errorf("error on decoding the blogs data from file")
		}
	}

	existingBlogs = append(existingBlogs, blogModel)
	jsonWriteData, err := json.Marshal(existingBlogs)
	if err != nil {
		return fmt.Errorf("error on writing to json")
	}
	err = os.WriteFile("data/store.json", jsonWriteData, filePermissionMode)
	if err != nil {
		return fmt.Errorf("error on writing the file")
	}
	return nil
}

// GetAll gets all the blogs based on pagination input.
func (blogService *JSONBlogService) GetAll(limit, offset int) []models.BlogModel {
	return []models.BlogModel{}
}

// GetBlogs gets the blog.
func (blogService *JSONBlogService) GetBlogs(id int) models.BlogModel {
	return models.BlogModel{}
}

// DeleteBlog deletes the blog.
func (blogService *JSONBlogService) DeleteBlog(id int) {

}

// GetRecentBlogs gets the recent blogs..
func (blogService *JSONBlogService) GetRecentBlogs() {

}

// InitJSONBlogService initializes the blog service.
func InitJSONBlogService() interfaces.Blog {
	b := new(JSONBlogService)
	return b
}
