package datastore

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type Redis struct {
	client *redis.Client
}

var (
	redisAddr     = "localhost:6379"
	redisPassword = ""
	redisDB       = "0"
)

func NewRedis() (*Redis, error) {
	a := getEnvVar("REDIS_ADDR", redisAddr)
	p := getEnvVar("REDIS_PASSWORD", redisPassword)
	db, err := strconv.Atoi(getEnvVar("REDIS_DB", redisDB))
	if err != nil {
		return nil, fmt.Errorf("Invalid value for Redis_DB: %v, error: %v\n", db, err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     a,
		Password: p,
		DB:       db,
	})
	if _, err = client.Ping().Result(); err != nil {
		return nil, fmt.Errorf("Failed connecting to Redis DB, error: %v\n", err)
	}
	return &Redis{
		client: client,
	}, nil
}

func (r *Redis) Get(url string) (string, bool, error) {
	val, err := r.client.Get(url).Result()
	if err == redis.Nil {
		return "", false, nil
	}
	if err != nil {
		return "", false, errors.Wrap(err, "Failed getting a value from Redis")
	}
	return val, true, nil
}

func getEnvVar(n, v string) string {
	if val := os.Getenv(n); val != "" {
		return val
	}
	return v
}
