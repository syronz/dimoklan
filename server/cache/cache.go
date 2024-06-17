package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"dimoklan/internal/config"
)

type Cache struct {
	core  config.Core
	redis *redis.Client
}

type Student struct {
	Name string `redis:"name"`
	Age  int    `redis:"age"`
}

func NewCache(core config.Core) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     core.GetRedisAddr(),     // Redis server address
		Password: core.GetRedisPassword(), // No password set
		DB:       core.GetRedisDB(),       // Use default DB
	})

	ctx := context.Background()
	err := client.Set(ctx, "foo2", "bar", 10*time.Second).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to save to the cahce; err: %w", err)
	}

	redisCache := &Cache{
		core:  core,
		redis: client,
	}

	student := Student{Name: "Adrian", Age: 3}

	// Set some fields.
	if _, err := client.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, "key", "name", student.Name)
		rdb.HSet(ctx, "key", "age", student.Age)
		return nil
	}); err != nil {
		panic(err)
	}

	return redisCache, nil
}
