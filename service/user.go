package service

import (
	"context"

	"github.com/DanielVieirass/um_help/config"
	"github.com/DanielVieirass/um_help/model"
	"github.com/DanielVieirass/um_help/repo"
	"github.com/rs/zerolog"
)

type UserService struct {
	Config *config.Config
	Logger *zerolog.Logger
	Repo   *repo.RepoManager
}

func NewUserService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *UserService {
	return &UserService{
		Config: cfg,
		Logger: logger,
		Repo:   repo,
	}
}

func (s *UserService) NewUser(ctx context.Context, r *req.NewUser) error {
	user := &model.User{
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		DocumentNumber: r.DocumentNumber,
		Balance:        0,
	}

	s.Logger.Info().Msgf("creating user %s", user.DocumentNumber)

	return s.Repo.MySQL.User.InsertUser(ctx, user)
}
