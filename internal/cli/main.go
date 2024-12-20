package cli

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
)

func Run(args []string) bool {
	// Загружаем конфигурацию
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Настраиваем логгер
	logger := config.SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	logger.Info("Starting gRPC and HTTP servers...")

	// Создаем контекст с уведомлением о прерывании
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Подключаемся к базе данных
	dbCon, err := data.NewDBConnection(cfg.Database.URL)
	if err != nil {
		logger.Fatalf("failed to connect to the database: %v", err)
	}
	defer func() {
		if err := dbCon.Close(); err != nil {
			logger.Fatalf("failed to close the database connection: %v", err)
		}
	}()

	// Создаем объект server
	service, err := config.NewServer(cfg)
	if err != nil {
		logger.Fatalf("failed to create server: %v", err)
		return false
	}

	ctx = cifractx.WithValue(ctx, config.SERVICE, service)

	var wg sync.WaitGroup
	runServices(ctx, &wg)

	wgch := make(chan struct{})
	go func() {
		wg.Wait()
		close(wgch)
	}()

	select {
	case <-ctx.Done():
		log.Printf("Interrupt signal received: %v", ctx.Err())
		<-wgch
	case <-wgch:
		log.Print("All services stopped")
	}

	return true
}
