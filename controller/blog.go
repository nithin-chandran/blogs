package controller

import (
	"blogs/interfaces"
	"blogs/models"
	"blogs/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BlogController - controller for the blogs.
type BlogController struct {
	blogService interfaces.Blog
}

// CreateBlog creates the blog.
func (blogController *BlogController) CreateBlog(c *gin.Context) {
	// get values from request
	reqBody := models.BlogRequest{}
	err := c.ShouldBind(&reqBody)
	if err != nil {
		log.Printf("bad request : %s \n", err.Error())
		c.JSON(400, "Invalid request")
		return
	}
	err = blogController.blogService.Create(models.BlogModel{
		ID:          reqBody.ID,
		Title:       reqBody.Title,
		Description: reqBody.Description,
	})
	if err != nil {
		log.Printf("error on creat: %s", err.Error())
		c.JSON(500, "internal server error")
		return
	}
	c.JSON(http.StatusOK, "success")
}

// CreateBlog creates the blog.
func (blogController *BlogController) GetBlogs() {

}

// InitBlogController initializes the blog controller.
func InitBlogController() *BlogController {
	b := new(BlogController)
	b.blogService = service.InitJSONBlogService()
	return b
}
