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
