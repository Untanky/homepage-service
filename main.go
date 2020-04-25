package main

import (
	"encoding/json"
	"fmt"
	"github.com/Kamva/mgm/v2"
	"github.com/gin-gonic/gin"
	"homepage-service/models"
	web "homepage-service/service"
	"net/http"
)

func main() {
	InitDB()

	post := models.Post{
		DefaultModel: mgm.DefaultModel{},
		Id:           0,
		Title:        "Test",
		Author:       "Lukas",
		Created:      0,
		Updated:      0,
		Section:      nil,
	}

	data, errors := json.Marshal(post)

	fmt.Println(string(data))
	fmt.Println(errors)

	service := gin.Default()

	api := service.Group("/api/v1")

	api.GET("/health", handleHealth)

	web.SetupPortfolioService(api.Group("/portfolio"))
	web.SetupBlogService(api.Group("/blog"))

	err := http.ListenAndServe(":8080", service)

	if err != nil {
		fmt.Println("Error creating web service")
	} else {
		fmt.Println("Webservice running on port 8080")
	}
}

func handleHealth(context *gin.Context) {
	context.Status(200)
}