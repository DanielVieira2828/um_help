package repo

import (
	"github.com/savi2w/pupper/config"
	"github.com/savi2w/pupper/repo/mysql"
	"github.com/savi2w/pupper/repo/redis"
)

type RepoManager struct {
	MySQL *mysql.Repo
	Redis *redis.Repo
}

func New(cfg *config.Config) (*RepoManager, error) {
	mysql, err := mysql.New(cfg)
	if err != nil {
		return nil, err
	}

	redis, err := redis.New(cfg)
	if err != nil {
		return nil, err
	}

	return &RepoManager{
		MySQL: mysql,
		Redis: redis,
	}, nil
}
