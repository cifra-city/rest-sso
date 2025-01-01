package config

import (
	"github.com/cifra-city/mailman"
	"github.com/cifra-city/rest-sso/internal/data/db"
	"github.com/cifra-city/tokens"
	"github.com/sirupsen/logrus"
)

const (
	SERVER = "service"
)

type Server struct {
	Config       *Config
	Databaser    *db.Databaser
	Logger       *logrus.Logger
	Mailman      *mailman.Mailman
	TokenManager *tokens.TokenManager
	Broker       *Broker
}

func NewServer(cfg *Config) (*Server, error) {
	logger := SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	mail := mailman.NewMailman(cfg.Email.SmtpPort, cfg.Email.SmtpHost, cfg.Email.Address, cfg.Email.Password)
	queries, err := db.NewDatabaser(cfg.Database.URL)
	tokenManager := tokens.NewTokenManager(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB, logger, cfg.JWT.AccessToken.TokenLifetime)
	broker, err := NewBroker(cfg.Rabbit.URL, cfg.Rabbit.Exchange)
	if err != nil {
		return nil, err
	}

	return &Server{
		Config:       cfg,
		Databaser:    queries,
		Logger:       logger,
		Mailman:      mail,
		TokenManager: tokenManager,
		Broker:       broker,
	}, nil
}
