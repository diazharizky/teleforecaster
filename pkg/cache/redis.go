package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Client *redis.Client
}

func New(host, port, password string, db int) *Cache {
	addr := fmt.Sprintf("%s:%s", host, port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := client.Ping(context.TODO()).Err(); err != nil {
		panic(err)
	}

	return &Cache{client}
}
