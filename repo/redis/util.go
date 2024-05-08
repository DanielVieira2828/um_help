package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Util struct {
	cli *redis.Client
}

func (u *Util) SetString(ctx context.Context, k string, v string, e time.Duration) error {
	if _, err := u.cli.Set(ctx, k, v, e).Result(); err != nil {
		return err
	}

	return nil
}

func (u *Util) GetString(ctx context.Context, k string) (string, error) {
	v, err := u.cli.Get(ctx, k).Result()

	if err != nil {
		if err == redis.Nil {
			return "", nil
		}

		return "", err
	}

	return v, nil
}
