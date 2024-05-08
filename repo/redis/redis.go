package redis

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"github.com/savi2w/pupper/config"
)

type Repo struct {
	Util *Util

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

	return &Repo{
		Util: &Util{cli: cli},

		cli: cli,
	}, nil
}
