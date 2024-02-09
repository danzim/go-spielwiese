package common

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	// Initialisiere den Logger
	log = logrus.New()

	// Setze das Log-Format auf JSON (kann je nach Bedarf angepasst werden)
	log.SetFormatter(&logrus.JSONFormatter{})
}

// GetLogger gibt den initialisierten Logger zurück
func GetLogger() *logrus.Logger {
	return log
}

// Example of how to use the logger in another file:
//
// import "path/to/internal/common"
//
// func someFunction() {
//     logger := common.GetLogger()
//     logger.Info("This is an info message")
//     logger.Error("This is an error message")
// }
