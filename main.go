package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	web "homepage-service/service"
	"net/http"
)

func main() {
	InitDB()

	service := gin.Default()

	api := service.Group("/api/v1")

	api.GET("/health", handleHealth)

	web.SetupPortfolioService(api.Group("/portfolio"))
	web.SetupProjectService(api.Group("/project"))

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