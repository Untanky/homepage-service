package service

import (
	"github.com/Kamva/mgm/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func getPosts(context *gin.Context) {
	raw := []bson.M{}

	err := mgm.CollectionByName("posts").SimpleFind(&raw, bson.M{})

	if err != nil {
		context.Status(500)
	}

	context.JSON(200, raw)
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
