package cli

import (
	"context"
	"sync"

	"github.com/cifra-city/rest-sso/internal/service"
)

func runServices(ctx context.Context, wg *sync.WaitGroup) {
	var (
	// signals indicate the finished initialization of each worker
	)

	run := func(f func()) {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}

	run(func() { service.Run(ctx) })
}
