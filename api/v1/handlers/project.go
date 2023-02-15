package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danzim/go-spielwiese/api/v1/models"
	"github.com/danzim/go-spielwiese/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func getRedis() (*db.GoRedis, error) {
	redisService, err := db.RedisConnect("127.0.0.1:6379")
	if err != nil {
		return nil, err
	}

	return redisService, nil
}

func GetProjects(c *gin.Context) {
	var projects []models.Project

	redisService, err := getRedis()
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	ctx := context.Background()

	log.Info().Msg("Get all CI Keys from Redis")
	keys := redisService.Client.Keys(ctx, "ci-*").Val()
	for _, s := range keys {
		//fmt.Println(i, s)
		val, err := redisService.Handler.JSONGet(s, ".")
		if err != nil {
			log.Error().Err(err).Msg("")
		}
		projectJSON := val.([]byte)
		readProject := models.Project{}

		err = json.Unmarshal(projectJSON, &readProject)
		if err != nil {
			fmt.Println(err)
		}
		projects = append(projects, readProject)
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
	//fmt.Println(projects)

}

func GetProject(c *gin.Context) {
	//var projects []models.Project

	redisService, err := getRedis()
	if err != nil {
		fmt.Println(err)
	}

	val, err := redisService.Handler.JSONGet(c.Param("id"), ".")
	if err != nil {
		fmt.Println(err)
	}
	projectJSON := val.([]byte)
	readProject := models.Project{}
	err = json.Unmarshal(projectJSON, &readProject)
	//er := enc.Encode(val)
	if err != nil {
		fmt.Println(err)
		//	panic("unable to encode value")
	}

	c.JSON(http.StatusOK, gin.H{"data": readProject})
}

func CreateUpdateProject(c *gin.Context) {
	var input models.Project

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	project := models.Project{Ci: c.Param("id"), DisplayName: input.DisplayName, Description: input.Description}

	redisService, err := getRedis()
	if err != nil {
		fmt.Println(err)
	}

	redisService.Handler.JSONSet(c.Param("id"), ".", project)

	c.JSON(http.StatusOK, gin.H{"data": project})
	//fmt.Println(project)
}

func DeleteProject(c *gin.Context) {

	redisService, err := getRedis()
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()

	key := redisService.Client.Keys(ctx, c.Param("id"))

	if key.Val()[0] == c.Param("id") {
		redisService.Handler.JSONDel(c.Param("id"), ".")
		c.JSON(http.StatusOK, gin.H{"data": true})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	//fmt.Println(key.Val()[0])
}
