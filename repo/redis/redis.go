package redis

import (
	"context"
	"time"

	"github.com/DanielVieirass/um_help/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

type Repo struct {
	cli *redis.Client
}

func New(cfg *config.Config) (*Repo, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisConfig.Host + ":" + cfg.RedisConfig.Port,
		Password: cfg.RedisConfig.Password,
		DB:       cfg.RedisConfig.Database,
	})

	ctx := context.Background()

	if _, err := cli.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &Repo{cli: cli}, nil
}

func (r *Repo) SetString(ctx context.Context, k string, v string, e time.Duration) error {
	if _, err := r.cli.Set(ctx, k, v, e).Result(); err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetString(ctx context.Context, k string) (string, error) {
	v, err := r.cli.Get(ctx, k).Result()

	if err != nil {
		if err == redis.Nil {
			return "", nil
		}

		return "", err
	}

	return v, nil
}
