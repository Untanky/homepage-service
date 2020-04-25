package service

import (
	"github.com/Kamva/mgm/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"homepage-service/models"
	"strconv"
)

func getPost(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 32)

	if err != nil {
		context.String(400, "ID must be int")
	}

	coll := mgm.Coll(&models.Post{})

	raw := models.Post{}
	err  = coll.First(bson.M{"id": id}, &raw)

	if err != nil {
		context.String(404, "Not found")
		return
	}

	context.JSON(200, raw)
}

func createPost(context *gin.Context) {
	context.Status(501)
}

func updatePost(context *gin.Context) {
	context.Status(501)
}

func deletePost(context *gin.Context) {
	context.Status(501)
}

func getPostPreviews(context *gin.Context) {
	context.Status(501)
}

func getPreviewOfPopularPosts(context *gin.Context) {
	context.Status(501)
}

func getPreviewOfNewPosts(context *gin.Context) {
	context.Status(501)
}

func SetupBlogService(api *gin.RouterGroup) {

	postApi := api.Group("/post")
	postApi.GET("/:id", getPost)
	postApi.POST("/", createPost)
	postApi.PUT("/:id", updatePost)
	postApi.DELETE("/:id", deletePost)

	previewApi := api.Group("/preview")
	previewApi.GET("/", getPostPreviews)
	previewApi.GET("/new", getPreviewOfNewPosts)
	previewApi.GET("/popular", getPreviewOfPopularPosts)
}
