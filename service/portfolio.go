package service

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path"
)

func getLanguagePathFromHeader(context *gin.Context) string {
	lang := context.GetHeader("Accept-Language")
	var langPath string

	switch lang {
	case "de":
		langPath = "de"
	case "en":
		fallthrough
	default:
		langPath = "en"
		break
	}

	return langPath
}

func getEducation(context *gin.Context) {
	langPath := getLanguagePathFromHeader(context)

	data, err := ioutil.ReadFile(path.Join("data", langPath, "education.json"))

	if err != nil {
		context.Status(500)
	}

	context.Data(200, "application/json", data)
}

func getExperience(context *gin.Context) {
	langPath := getLanguagePathFromHeader(context)

	data, err := ioutil.ReadFile(path.Join("data", langPath,"experience.json"))

	if err != nil {
		context.Status(500)
	}

	context.Data(200, "application/json", data)
}

func getSkills(context *gin.Context) {
	langPath := getLanguagePathFromHeader(context)

	data, err := ioutil.ReadFile(path.Join("data", langPath, "skills.json"))

	if err != nil {
		context.Status(500)
	}

	context.Data(200, "application/json", data)
}

func getCategories(context *gin.Context) {
	langPath := getLanguagePathFromHeader(context)

	data, err := ioutil.ReadFile(path.Join("data", langPath, "categories.json"))

	if err != nil {
		context.Status(500)
	}

	context.Data(200, "application/json", data)
}

func getStrengths(context *gin.Context) {
	langPath := getLanguagePathFromHeader(context)

	data, err := ioutil.ReadFile(path.Join("data", langPath,"strengths.json"))

	if err != nil {
		context.Status(500)
	}

	context.Data(200, "application/json", data)
}

func getLanguages(context *gin.Context) {
	langPath := getLanguagePathFromHeader(context)

	data, err := ioutil.ReadFile(path.Join("data", langPath,"languages.json"))

	if err != nil {
		context.Status(500)
	}

	context.Data(200, "application/json", data)
}

func getCV(context *gin.Context) {
	langPath := getLanguagePathFromHeader(context)

	data, err := ioutil.ReadFile(path.Join("data", langPath,"cv.pdf"))

	if err != nil {
		context.Status(500)
	}

	context.Data(200, "application/pdf", data)
}

func SetupPortfolioService(api *gin.RouterGroup) {
	api.GET("/education", getEducation)
	api.GET("/experience", getExperience)
	api.GET("/skills", getSkills)
	api.GET("/categories", getCategories)
	api.GET("/strengths", getStrengths)
	api.GET("/languages", getLanguages)
	api.GET("/cv", getCV)
}