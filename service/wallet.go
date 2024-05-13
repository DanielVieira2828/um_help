package service

import (
	"context"

	"github.com/DanielVieirass/um_help/config"
	"github.com/DanielVieirass/um_help/model"
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

func (s *WalletService) New(ctx context.Context, currencyId int64, ownerId int64, alias string) error {
	wallet := &model.Wallet{
		OwnerId:    ownerId,
		Alias:      alias,
		CurrencyId: currencyId,
	}

	s.logger.Info().Msgf("creating wallet for user %s", wallet.OwnerId)

	return s.repo.MySQL.Wallet.InsertWallet(ctx, wallet)
}
