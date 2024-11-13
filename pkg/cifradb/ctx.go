package cifradb

import (
	"context"
	"fmt"

	"github.com/cifra-city/rest-sso/internal/db/data"
)

type contextKey string

const dbQueriesKey contextKey = "dbQueries"

// WithDBQueries добавляет объект queries в контекст.
func WithDBQueries(ctx context.Context, queries *data.Queries) context.Context {
	return context.WithValue(ctx, dbQueriesKey, queries)
}

// GetDBQueries получает объект queries из контекста.
func GetDBQueries(ctx context.Context) (*data.Queries, error) {
	queries, ok := ctx.Value(dbQueriesKey).(*data.Queries)
	if !ok {
		return nil, fmt.Errorf("could not get queries from context")
	}
	return queries, nil
}
