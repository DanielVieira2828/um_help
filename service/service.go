package service

import (
	"github.com/DanielVieirass/um_help/config"
	"github.com/DanielVieirass/um_help/repo"
	"github.com/DanielVieirass/um_help/util/cryptoutil"
	"github.com/rs/zerolog"
)

type Service struct {
	Auth   *AuthService
	User   *UserService
	Wallet *WalletService
}

func New(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) (*Service, error) {
	cryptoutil, err := cryptoutil.New(cfg)
	if err != nil {
		return nil, err
	}

	return &Service{
		Auth:   newAuthService(cfg, cryptoutil, logger, repo),
		User:   newUserService(cfg, cryptoutil, logger, repo),
		Wallet: newWalletService(cfg, logger, repo),
	}, nil
}
