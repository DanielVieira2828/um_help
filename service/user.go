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
	config *config.Config
	logger *zerolog.Logger
	repo   *repo.RepoManager
}

func newUserService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *UserService {
	return &UserService{
		config: cfg,
		logger: logger,
		repo:   repo,
	}
}

func (s *UserService) New(ctx context.Context, r *req.NewUser) error {
	user := &model.User{
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		DocumentNumber: r.DocumentNumber,
		Balance:        10000,
	}

	s.logger.Info().Msgf("creating user %s", user.DocumentNumber)

	return s.repo.MySQL.User.InsertUser(ctx, user)
}
