package service

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path"
)

func getProjects(context *gin.Context) {
	langPath := getLanguagePathFromHeader(context)

	data, err := ioutil.ReadFile(path.Join("data", langPath, "projects.json"))

	if err != nil {
		context.Status(500)
	}

	context.Data(200, "application/json", data)
}

func SetupProjectService(api *gin.RouterGroup) {
	api.GET("/", getProjects)
}
