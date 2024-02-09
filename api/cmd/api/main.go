package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/danzim/go-spielwiese/v2/internal/schemas"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func main() {

	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// ClusterService Endpunkte
	r.GET("/clusters", listClusters)
	r.GET("/clusters/:cluster", readCluster)

	// ProjectService Endpunkte
	r.GET("/clusters/:cluster/projects", listProjectsInCluster)
	r.POST("/clusters/:cluster/projects", createProject)
	r.GET("/clusters/:cluster/projects/:project", readProject)
	r.PATCH("/clusters/:cluster/projects/:project", updateProject)
	r.PUT("/clusters/:cluster/projects/:project", replaceProject)
	r.DELETE("/clusters/:cluster/projects/:project", deleteProject)

	// QuayService Endpunkte
	r.PUT("/quay/organizations/:organization", reconfigureQuayOrganization)

	r.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:9080")
}

func loadConfig() (*schemas.APIconfig, error) {
	configFile := "config/api_config.yaml"
	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	config := &schemas.APIconfig{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("failed to decode config: %v", err)
	}

	return config, nil
}

// Funktionen für die Endpunkte (Handler-Funktionen)
func listClusters(c *gin.Context) {
	// Hier kannst du die Logik für die "listClusters"-Funktion implementieren
}

func readCluster(c *gin.Context) {
	// Hier kannst du die Logik für die "readCluster"-Funktion implementieren
}

func listProjectsInCluster(c *gin.Context) {
	// Hier kannst du die Logik für die "listProjectsInCluster"-Funktion implementieren
}

func createProject(c *gin.Context) {
	// Hier kannst du die Logik für die "createProject"-Funktion implementieren
}

func readProject(c *gin.Context) {
	// Hier kannst du die Logik für die "createProject"-Funktion implementieren
}

func updateProject(c *gin.Context) {
	// Hier kannst du die Logik für die "createProject"-Funktion implementieren
}

func replaceProject(c *gin.Context) {
	// Hier kannst du die Logik für die "createProject"-Funktion implementieren
}

func deleteProject(c *gin.Context) {
	// Hier kannst du die Logik für die "createProject"-Funktion implementieren
}

func reconfigureQuayOrganization(c *gin.Context) {
	// Hier kannst du die Logik für die "createProject"-Funktion implementieren
}
