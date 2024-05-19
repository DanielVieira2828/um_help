package service

import (
	"github.com/DanielVieirass/um_help/config"
	"github.com/DanielVieirass/um_help/repo"
	"github.com/rs/zerolog"
)

type AuthService struct {
	config *config.Config
	logger *zerolog.Logger
	repo   *repo.RepoManager
}

func newAuthService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *AuthService {
	return &AuthService{
		config: cfg,
		logger: logger,
		repo:   repo,
	}
}
