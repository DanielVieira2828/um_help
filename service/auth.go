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

func newAuthService(cfg *config.Config, cryptoutil *cryptoutil.Cryptoutil, logger *zerolog.Logger, repo *repo.RepoManager) *AuthService {
	return &AuthService{
		config:     cfg,
		cryptoutil: cryptoutil,
		logger:     logger,
		repo:       repo,
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

	if s.cryptoutil.HashString(r.Password) != user.Password {
		return nil, errors.New("wrong credentials")
	}

	result, err := s.cryptoutil.SignUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err = s.repo.Redis.SetStruct(ctx, result.SignId, result, 0); err != nil {
		return nil, err
	}

	resp := &res.LoginResponse{
		JWS:            result.JWS,
		ExpirationTime: result.ExpirationTime,
		RefreshToken:   result.RefreshToken,
	}

	return resp, nil
}
