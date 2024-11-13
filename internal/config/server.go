package config

import (
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/sirupsen/logrus"
)

const (
	SERVICE = "service"
)

type Service struct {
	Config  *Config
	Queries *data.Queries
	Logger  *logrus.Logger
}

func NewServer(cfg *Config) (*Service, error) {
	// Настройка логгера
	logger := SetupLogger(cfg.Logging.Level, cfg.Logging.Format)

	// Подключение к базе данных
	dbCon, err := data.NewDBConnection(cfg.Database.URL)
	if err != nil {
		logger.Fatalf("failed to connect to the database: %v", err)
		return nil, err
	}

	// Создание queries
	queries := data.New(dbCon)

	return &Service{
		Config:  cfg,
		Queries: queries,
		Logger:  logger,
	}, nil
}
