package httpresp

import (
	"context"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

// StartServer запускает HTTP сервер в отдельной горутине и возвращает указатель на сервер.
func StartServer(ctx context.Context, addr string, router http.Handler) *http.Server {
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		logrus.Infof("Starting server on port %s", addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("Server failed to start: %v", err)
		}
	}()

	return server
}

// StopServer завершает работу сервера при отмене контекста.
func StopServer(ctx context.Context, server *http.Server) {
	logrus.Info("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		logrus.Errorf("Server shutdown failed: %v", err)
	}
}
