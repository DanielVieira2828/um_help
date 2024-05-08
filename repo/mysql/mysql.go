package mysql

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/savi2w/pupper/config"
)

type Repo struct {
	cli *sqlx.DB
}

func New(cfg *config.Config) (*Repo, error) {
	url := cfg.MySQLConfig.Username + ":" + cfg.MySQLConfig.Password + "@tcp(" + cfg.MySQLConfig.Host + ":" + cfg.MySQLConfig.Port + ")/" + cfg.MySQLConfig.Database + "?parseTime=true"

	cli, err := sqlx.Connect("mysql", url)
	if err != nil {
		return nil, err
	}

	cli.DB.SetConnMaxLifetime(time.Minute * 5)
	cli.DB.SetMaxIdleConns(5)
	cli.DB.SetMaxOpenConns(100)

	if err := cli.Ping(); err != nil {
		return nil, err
	}

	return &Repo{
		cli: cli,
	}, nil
}
