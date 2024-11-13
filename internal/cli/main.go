package cli

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/rest-sso/pkg/cifradb"
)

func Run(args []string) bool {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := config.SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	logger.Info("Starting gRPC and HTTP servers...")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	dbCon, err := data.NewDBConnection(cfg.Database.URL)
	if err != nil {
		logger.Fatalf("failed to connect to the database: %v", err)
	}
	defer func(dbCon *sql.DB) {
		err := dbCon.Close()
		if err != nil {
			logger.Fatalf("failed to close the database connection: %v", err)
		}
	}(dbCon)

	queries := data.New(dbCon)
	ctx = cifradb.WithDBQueries(ctx, queries)

	var wg sync.WaitGroup

	runServices(ctx, *cfg, &wg)

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

	wgch := make(chan struct{})
	go func() {
		wg.Wait()
		close(wgch)
	}()

	select {
	case <-ctx.Done():
		log.Printf("Interrupt signal received %s", ctx.Err())
		<-wgch
	case <-wgch:
		log.Print("all services stopped")
	}

	return true
}
