package service

import (
	"github.com/Kamva/mgm/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"homepage-service/models"
	"homepage-service/url"
	"math"
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

func viewPost(context *gin.Context) {
	context.Status(501)
}

func lovePost(context *gin.Context) {
	context.Status(501)
}

//func getPostPreviews(context *gin.Context) {
//	coll := mgm.Coll(&models.Post{})
//
//	ctx := mgm.Ctx()
//	data, err := coll.Find(ctx, bson.M{})
//
//	if err != nil {
//		context.String(500, err.Error())
//		return
//	}
//
//	raw := []models.Post{}
//	err = data.All(ctx, &raw)
//
//	if err != nil {
//		context.String(404, "Not found")
//		return
//	}
//
//	var previews []models.Preview
//
//	for i := range raw {
//		previews = append(previews, raw[i].GeneratePreview())
//	}
//
//	context.JSON(200, previews)
//}

const pageSize = 6

func pageData(context *gin.Context) (int, int) {
	var first = int64(0)

	if start, _ := context.GetQuery("start"); start != "" {
		if num, err := strconv.ParseInt(start, 10, 32); err == nil {
			first = num
		}
	}

	var last = first + pageSize

	if end, _ := context.GetQuery("end"); end != "" {
		if num, err := strconv.ParseInt(end, 10, 32); err == nil {
			last = num
		}

		if last - first > pageSize {
			last = first + pageSize
		}
	}

	return int(first), int(last)
}

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
		return raw[i].Created.Time().After(raw[j].Created.Time())
	}

	sort.Slice(raw, byDate)

	var previews []models.Preview

	for i := range raw {
		previews = append(previews, raw[i].GeneratePreview())
	}

	first, last := pageData(context)

	if first < 0 {
		first = 0
	}

	if last >= len(previews) {
		last = len(previews)
	}

	if last < first {
		context.String(400, "Wrong pagination")
	}

	context.JSON(200, previews[first:last])
}

func getPreviewOfPopularPosts(context *gin.Context) {
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

	byPopularity := func(i, j int) bool {
		return raw[i].GetScore() > raw[j].GetScore()
	}

	sort.Slice(raw, byPopularity)

	var previews []models.Preview

	for i := range raw {
		previews = append(previews, raw[i].GeneratePreview())
	}

	first, last := pageData(context)

	if first < 0 {
		first = 0
	}

	if last >= len(previews) {
		last = len(previews)
	}

	if last < first {
		context.String(400, "Wrong pagination")
	}

	context.JSON(200, previews[first:last])
}

func getPreviewsForPost(context *gin.Context) {
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

	originalPostId, err := strconv.ParseInt(context.Param("id"), 10, 32)

	if err != nil {
		context.String(400, "Please provide integer as id")
		return
	}

	byPost := func(i, j int) bool {
		return math.Abs(float64(raw[i].Id - int(originalPostId))) < math.Abs(float64(raw[j].Id - int(originalPostId)))
	}

	sort.Slice(raw, byPost)

	raw = raw[1:]

	var previews []models.Preview

	for i := range raw {
		previews = append(previews, raw[i].GeneratePreview())
	}

	first, last := pageData(context)

	if first < 0 {
		first = 0
	}

	if last >= len(previews) {
		last = len(previews)
	}

	if last < first {
		context.String(400, "Wrong pagination")
	}

	context.JSON(200, previews[first:last])
}

func SetupBlogService(api *gin.RouterGroup) {

	postApi := api.Group("/post")
	postApi.GET("/:id", getPost)
	postApi.POST("/", createPost)
	postApi.PUT("/:id", updatePost)
	postApi.DELETE("/:id", deletePost)

	postApi.POST("/:id/view", viewPost)
	postApi.POST("/:id/love", lovePost)

	previewApi := api.Group("/preview")
	//previewApi.GET("/", getPostPreviews)
	previewApi.GET("/new", getPreviewOfNewPosts)
	previewApi.GET("/popular", getPreviewOfPopularPosts)
	previewApi.GET("/for/:id", getPreviewsForPost)

	url.PostPath = postApi.BasePath()
	url.PreviewPath = previewApi.BasePath()
}
