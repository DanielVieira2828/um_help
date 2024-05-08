package service

import (
	"github.com/rs/zerolog"
	"github.com/savi2w/pupper/config"
	"github.com/savi2w/pupper/repo"
)

type Service struct {
}

func New(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *Service {
	return &Service{}
}
