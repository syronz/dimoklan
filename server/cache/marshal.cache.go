package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"dimoklan/model"
)

func (c *Cache) AddMarshalMoveToFraction(ctx context.Context, fraction string, moveMarshal model.MarshalMove) error {
	_, err := c.redis.HSet(ctx, fraction, moveMarshal.MarshalID, moveMarshal.ToZipString()).Result()
	return err
}

func (c *Cache) AddMarshalMove(ctx context.Context, moveMarshal model.MarshalMove) error {
	_, err := c.redis.HSet(ctx, moveMarshal.MarshalID, moveMarshal).Result()

	return err
}

func (c *Cache) GetMarshalMove(ctx context.Context, marshalID string) (moveMarshal model.MarshalMove, err error) {
	err = c.redis.HGetAll(ctx, marshalID).Scan(&moveMarshal)
	fmt.Printf(">>>>>>> old12: %+v\n", err)

	return moveMarshal, err
}

func (c *Cache) SaveMove(ctx context.Context, moveMarshal model.MarshalMove) error {
	/*
		var ongoingMove model.MoveMarshal

		if err := c.redis.HGetAll(ctx, hashtag.Marshal+marshalID).Scan(&ongoingMove); err != nil {
			return err
		}

		dest := localtype.CELL(ongoingMove.Destination)




		_, err := c.redis.Pipelined(ctx, func(redis.Pipeliner) error {

		})

		if err != nil {
			return err
		}

		// if _, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		// 	rdb.HSet(ctx, "key", "str1", "hello")
		// 	rdb.HSet(ctx, "key", "str2", "world")
		// 	rdb.HSet(ctx, "key", "int", 123)
		// 	rdb.HSet(ctx, "key", "bool", 1)
		// 	rdb.HSet(ctx, "key", "bytes", []byte("this is bytes !"))
		// 	return nil
		// }); err != nil {
		// 	panic(err)
		// }
	*/
	return nil
}

func (c *Cache) DeleteMarshalMove(ctx context.Context, marshalID, sourceFraction, destFraction string) error {
	_, err := c.redis.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		c.redis.Del(ctx, marshalID)
		c.redis.HDel(ctx, sourceFraction, marshalID)
		c.redis.HDel(ctx, destFraction, marshalID)
		return nil
	})

	return err
}
