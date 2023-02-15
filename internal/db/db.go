package db

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
)

type GoRedis struct {
	Client  *redis.Client
	Handler *rejson.Handler
}

type Test struct {
	Message string `json:"message"`
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func RedisConnect(address string) (*GoRedis, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}

	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(client)

	return &GoRedis{
		Client:  client,
		Handler: rh,
	}, nil
}
