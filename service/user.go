package service

import (
	"context"

	"github.com/DanielVieirass/um_help/config"
	"github.com/DanielVieirass/um_help/model"
	"github.com/DanielVieirass/um_help/presenter/req"
	"github.com/DanielVieirass/um_help/repo"
	"github.com/rs/zerolog"
)

type UserService struct {
	config        *config.Config
	logger        *zerolog.Logger
	repo          *repo.RepoManager
	walletService *WalletService
}

func newUserService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager, walletService *WalletService) *UserService {
	return &UserService{
		config:        cfg,
		logger:        logger,
		repo:          repo,
		walletService: walletService,
	}
}

const BRL = 1

func (s *UserService) New(ctx context.Context, r *req.NewUser) error {
	user := &model.User{
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		DocumentNumber: r.DocumentNumber,
	}

	s.logger.Info().Msgf("creating user %s", user.DocumentNumber)

	userID, err := s.repo.MySQL.User.InsertUser(ctx, user)
	if err != nil {
		return err
	}

	err = s.walletService.New(ctx, userID, BRL, "Default Wallet")
	if err != nil {
		return err
	}

	return nil
}
