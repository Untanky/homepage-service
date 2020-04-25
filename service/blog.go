package service

import (
	"fmt"
	"github.com/Kamva/mgm/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"homepage-service/models"
	"homepage-service/url"
	"sort"
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
	coll := mgm.Coll(&models.Post{})

	ctx := mgm.Ctx()
	data, err := coll.Find(ctx, bson.M{})

	if err != nil {
		context.String(500, err.Error())
		return
	}

	raw := []models.Post{}
	err = data.All(ctx, &raw)

	if err != nil {
		context.String(404, "Not found")
		return
	}

	var previews []models.Preview

	for i := range raw {
		previews = append(previews, raw[i].GeneratePreview())
	}

	context.JSON(200, previews)
}

const pageSize = 3

func getPreviewOfNewPosts(context *gin.Context) {
	coll := mgm.Coll(&models.Post{})

	ctx := mgm.Ctx()
	data, err := coll.Find(ctx, bson.M{})

	if err != nil {
		context.String(500, err.Error())
		return
	}

	raw := []models.Post{}
	err = data.All(ctx, &raw)

	if err != nil {
		context.String(404, "Not found")
		return
	}

	byDate := func(i, j int) bool {
		return raw[i].Created.Time().Before(raw[j].Created.Time())
	}

	sort.Slice(raw, byDate)

	var previews []models.Preview

	for i := range raw {
		previews = append(previews, raw[i].GeneratePreview())
	}

	context.JSON(200, previews)
}

func getPreviewOfPopularPosts(context *gin.Context) {
	context.Status(501)
}

func getPreviewsForPost(context *gin.Context) {
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
	previewApi.GET("/for/:id", getPreviewsForPost)
	previewApi.GET("/test", test)

	url.PostPath = postApi.BasePath()
	url.PreviewPath = previewApi.BasePath()
}

func test(context *gin.Context) {
	var first, last int64

	if start, _ := context.GetQuery("start"); start != "" {
		if num, err := strconv.ParseInt(start, 10, 32); err != nil {
			first = 0
		} else {
			first = num
		}
	} else {
		first = 0
	}

	if end, _ := context.GetQuery("end"); end != "" {
		if num, err := strconv.ParseInt(end, 10, 32); err != nil {
			last = first + pageSize
		} else {
			last = num
		}

		if last - first > pageSize {
			last = first + pageSize
		}
	} else {
		last = first + pageSize
	}

	context.String(200, fmt.Sprintf("%d - %d", first, last))
}
