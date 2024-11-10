package cli

import (
	"context"
	"log"
	"sync"

	"github.com/cifra-city/rest-sso/internal/config"
)

func Run(args []string) bool {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	ctx := context.Background()
	var wg sync.WaitGroup
	runServices(ctx, *cfg, &wg)

	return true
}
