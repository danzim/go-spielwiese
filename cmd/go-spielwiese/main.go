package main

import (
	"net/http"

	"github.com/danzim/go-spielwiese/api/v1/handlers"
	"github.com/danzim/go-spielwiese/api/v1/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	//r := gin.Default()

	r := gin.New()
	r.Use(middleware.DefaultStructuredLogger())
	r.Use(gin.Recovery())

	//Misc
	r.GET("/version", handlers.GetVersion)
	r.POST("/version", handlers.SetVersion)

	//Project
	r.GET("projects", handlers.GetProjects)
	r.GET("/projects/:id", handlers.GetProject)
	r.POST("/projects/:id", handlers.CreateUpdateProject)
	r.PATCH("/projects/:id", handlers.CreateUpdateProject)
	r.DELETE("projects/:id", handlers.DeleteProject)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run()

}
