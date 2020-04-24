package service

import (
	"github.com/gin-gonic/gin"
)

func getPosts(context *gin.Context) {
	context.Status(401)
}

func getPost(context *gin.Context) {
	context.Status(401)
}

func createPost(context *gin.Context) {
	context.Status(401)
}

func updatePost(context *gin.Context) {
	context.Status(401)
}

func deletePost(context *gin.Context) {
	context.Status(401)
}

func SetupBlogService(api *gin.RouterGroup) {
	api.GET("/", getPosts)
	api.GET("/:id", getPost)
	api.POST("/", createPost)
	api.PUT("/:id", updatePost)
	api.DELETE("/:id", deletePost)
}
