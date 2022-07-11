package server

import (
	"blogs/controller"

	"blogs/config"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	config.InitConfig()

	blogController := controller.InitBlogController()
	router.POST("/blogs/", blogController.CreateBlog)
	router.Run()
}
