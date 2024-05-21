package service

import (
	"context"
	"errors"

	"github.com/DanielVieirass/um_help/config"
	"github.com/DanielVieirass/um_help/presenter/req"
	"github.com/DanielVieirass/um_help/presenter/res"
	"github.com/DanielVieirass/um_help/repo"
	"github.com/DanielVieirass/um_help/util/cryptoutil"
	"github.com/rs/zerolog"
)

type AuthService struct {
	config     *config.Config
	cryptoutil *cryptoutil.Cryptoutil
	logger     *zerolog.Logger
	repo       *repo.RepoManager
}

func newAuthService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *AuthService {
	return &AuthService{
		config: cfg,
		logger: logger,
		repo:   repo,
	}
}

func (s *AuthService) Login(ctx context.Context, r *req.LoginRequest) (*res.LoginResponse, error) {
	user, found, err := s.repo.MySQL.User.SelectByDocumentNumber(nil, ctx, r.DocumentNumber)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("user not found")
	}

	if s.cryptoutil.HashPassword(r.Password) != user.Password {
		return nil, errors.New("wrong credentials")
	}

	jws, expirationTime, err := s.cryptoutil.SignUserID(user.Id)
	if err != nil {
		return nil, err
	}

	resp := &res.LoginResponse{
		JWS:            jws,
		ExpirationTime: expirationTime,
	}

	return resp, nil
}
