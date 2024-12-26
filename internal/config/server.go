package config

import (
	"github.com/cifra-city/mailman"
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/tokens"
	"github.com/sirupsen/logrus"
)

const (
	SERVICE = "service"
)

type Service struct {
	Config       *Config
	Databaser    *data.Databaser
	Logger       *logrus.Logger
	Mailman      *mailman.Mailman
	TokenManager *tokens.TokenManager
}

func NewServer(cfg *Config) (*Service, error) {
	logger := SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	mail := mailman.NewMailman(cfg.Email.SmtpPort, cfg.Email.SmtpHost, cfg.Email.Address, cfg.Email.Password)
	queries, err := data.NewDatabaser(cfg.Database.URL)
	TokenManager := tokens.NewTokenManager(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB, cfg.JWT.AccessToken.TokenLifetime)

	if err != nil {
		return nil, err
	}

	return &Service{
		Config:       cfg,
		Databaser:    queries,
		Logger:       logger,
		Mailman:      mail,
		TokenManager: TokenManager,
	}, nil
}
