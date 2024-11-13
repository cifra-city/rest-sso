package cifradb

import (
	"net/http"

	"github.com/cifra-city/rest-sso/internal/db/data"
)

// DBQueriesMiddleware добавляет объект queries в контекст каждого запроса.
// Используйте его как middleware для chi.Router.
func DBQueriesMiddleware(queries *data.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Создаем новый контекст запроса с queries и передаем дальше
			reqCtx := WithDBQueries(r.Context(), queries)
			next.ServeHTTP(w, r.WithContext(reqCtx))
		})
	}
}
