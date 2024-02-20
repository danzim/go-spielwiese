package handlers

import (
	"net/http"

	"github.com/danzim/go-spielwiese/v2/internal/schemas"
	"github.com/gin-gonic/gin"
)

// ClusterHandler enthält Funktionen für Clusteroperationen
type ClusterHandler struct {
	redisHandler *RedisHandler
}

// NewClusterHandler erstellt eine neue Instanz des ClusterHandlers
func NewClusterHandler(redisHandler *RedisHandler) *ClusterHandler {
	return &ClusterHandler{
		redisHandler: redisHandler,
	}
}

// Funktionen für die Endpunkte (Handler-Funktionen)
func (ch *ClusterHandler) ListClusters(c *gin.Context) {

	var clusters schemas.ClusterList

	err := ch.redisHandler.GetData("clusters", &clusters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get clusters in Redis"})
		return
	}

	c.JSON(http.StatusOK, clusters)
}
