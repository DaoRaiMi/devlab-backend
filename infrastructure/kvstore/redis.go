package kvstore

import (
	redis "github.com/go-redis/redis"
)

type Redis struct {
	client redis.Client
}

func NewRedis(address, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
}
