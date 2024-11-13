package config

import (
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/sirupsen/logrus"
)

const (
	SERVER = "server"
)

type Server struct {
	Config  *Config
	Queries *data.Queries
	Logger  *logrus.Logger
}

func NewServer(cfg *Config) (*Server, error) {
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

	return &Server{
		Config:  cfg,
		Queries: queries,
		Logger:  logger,
	}, nil
}
