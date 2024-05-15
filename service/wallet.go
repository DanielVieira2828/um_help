package service

import (
	"github.com/DanielVieirass/um_help/config"
	"github.com/DanielVieirass/um_help/repo"
	"github.com/rs/zerolog"
)

type WalletService struct {
	config *config.Config
	logger *zerolog.Logger
	repo   *repo.RepoManager
}

func newWalletService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *WalletService {
	return &WalletService{
		config: cfg,
		logger: logger,
		repo:   repo,
	}
}
