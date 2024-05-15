package service

import (
	"context"
	"fmt"
	"strings"

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

func (s *UserService) New(ctx context.Context, r *req.NewUser) (u *model.User, err error) {
	user := &model.User{
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		DocumentNumber: r.DocumentNumber,
	}

	tx, err := s.repo.MySQL.BeginReadCommittedTx(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	userId, err := s.repo.MySQL.User.Insert(tx, ctx, user)
	if err != nil {
		return nil, err
	}

	user.Id = userId

	currency, found, err := s.repo.MySQL.Currency.SelectByCurrencyCode(tx, ctx, model.CurrencyBRL)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("cannot find `%s` currency in database", model.CurrencyBRL)
	}

	w := &model.Wallet{
		OwnerId:    user.Id,
		CurrencyId: currency.Id,
		Alias:      strings.Join([]string{user.FirstName + "'s", "Wallet"}, " "),
	}

	if err := s.repo.MySQL.Wallet.Insert(tx, ctx, w); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}
