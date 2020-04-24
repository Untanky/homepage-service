package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func SetupPortfolioService(api *gin.RouterGroup) {
	api.GET("/education", getEducation)
	api.GET("/experience", getExperience)
	api.GET("/skills", getSkills)
	api.GET("/categories", getCategories)
	api.GET("/strengths", getStrengths)
}

func getEducation(context *gin.Context) {
	data, err := ioutil.ReadFile("data/education.json")

	if err != nil {
		context.Status(500)
	}

	context.Data(200, "application/json", data)
}

func getExperience(context *gin.Context) {
	data, err := ioutil.ReadFile("data/experience.json")

	fmt.Println(data)

	if err != nil {
		context.Status(500)
	}

	context.Data(200, "application/json", data)
}

func getSkills(context *gin.Context) {
	data, err := ioutil.ReadFile("data/skills.json")

	if err != nil {
		context.Status(500)
	}

	context.Data(200, "application/json", data)
}

func getCategories(context *gin.Context) {
	data, err := ioutil.ReadFile("data/categories.json")

	if err != nil {
		context.Status(500)
	}

	context.Data(200, "application/json", data)
}

func getStrengths(context *gin.Context) {
	data, err := ioutil.ReadFile("data/strengths.json")

	if err != nil {
		context.Status(500)
	}

	context.Data(200, "application/json", data)
}