package service

import (
	"context"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/service/handlers"
	"github.com/cifra-city/rest-sso/pkg/cifradb"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
)

func Run(ctx context.Context, cfg config.Config) {
	r := chi.NewRouter()

	queries, err := cifradb.GetDBQueries(ctx)
	if err != nil {
		logrus.Fatalf("Failed to get DB queries from context: %v", err)
	}

	// Middleware для добавления queries в контекст запроса
	r.Use(cifradb.DBQueriesMiddleware(queries))

	// Определение маршрутов
	r.Route("/cifra-sso", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/public", func(r chi.Router) {
				r.Route("/auth", func(r chi.Router) {
					r.Post("/register", handlers.Registration)
				})
			})
		})
	})

	// Запуск сервера
	server := httpresp.StartServer(ctx, ":8080", r)

	// Ожидание отмены контекста для остановки сервера
	<-ctx.Done()
	httpresp.StopServer(context.Background(), server)
}
