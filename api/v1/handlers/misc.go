package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danzim/go-spielwiese/api/v1/models"
	"github.com/danzim/go-spielwiese/internal/db"
	"github.com/gin-gonic/gin"
)

func SetVersion(c *gin.Context) {
	version := models.Version{Version: "1.0.0"}

	redisService, err := db.RedisConnect("127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to Redis failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Connect to Redis failed"})
	}

	res, err := redisService.Handler.JSONSet("version", ".", version)
	if err != nil {
		fmt.Println("Failed to JSONSet with go redis")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to JSONSet with go redis"})
	}

	if res.(string) == "OK" {
		fmt.Printf("GoRedis Success: %s\n", res)
		c.JSON(http.StatusOK, version)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to JSONSet with go redis"})
	}

}

func GetVersion(c *gin.Context) {
	//var buf bytes.Buffer
	//enc := gob.NewEncoder(&buf)

	redisService, err := db.RedisConnect("127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
	}

	val, err := redisService.Handler.JSONGet("version", ".")
	if err != nil {
		fmt.Println(err)
	}

	versionJSON := val.([]byte)

	readVersion := models.Version{}

	err = json.Unmarshal(versionJSON, &readVersion)
	//er := enc.Encode(val)
	if err != nil {
		fmt.Println(err)
		//	panic("unable to encode value")
	}

	c.JSON(http.StatusOK, gin.H{"data": readVersion})

}
