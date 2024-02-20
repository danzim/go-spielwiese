package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisHandler enthält die Funktionalitäten für die Kommunikation mit Redis
type RedisHandler struct {
	client *redis.Client
}

// NewRedisHandler erstellt eine neue Instanz des Redis-Handlers
func NewRedisHandler(address, password string) (*RedisHandler, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0, // Standard-Datenbank
	})

	// Überprüfen, ob die Verbindung zu Redis erfolgreich hergestellt wurde
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return &RedisHandler{client: client}, nil
}

// SetData speichert JSON-Daten in Redis
func (rh *RedisHandler) SetData(key string, data interface{}, expiration time.Duration) error {
	ctx := context.Background()

	// Daten in JSON konvertieren
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data to JSON: %v", err)
	}

	// Daten in Redis speichern
	err = rh.client.Set(ctx, key, jsonData, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set data in Redis: %v", err)
	}

	return nil
}

// GetData liest JSON-Daten aus Redis
func (rh *RedisHandler) GetData(key string, dest interface{}) error {
	ctx := context.Background()

	// Daten aus Redis lesen
	result, err := rh.client.Get(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to get data from Redis: %v", err)
	}

	// JSON-Daten in die Zielschnittstelle konvertieren
	err = json.Unmarshal([]byte(result), dest)
	if err != nil {
		return fmt.Errorf("failed to unmarshal data from JSON: %v", err)
	}

	return nil
}
