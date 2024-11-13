package cifractx

import (
	"context"
	"fmt"
)

type ContextKey string

// WithValue добавляет произвольное значение в контекст с заданным ключом.
func WithValue(ctx context.Context, key ContextKey, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

// GetValue извлекает значение из контекста по заданному ключу и приводит его к нужному типу.
func GetValue[T any](ctx context.Context, key ContextKey) (T, error) {
	val, ok := ctx.Value(key).(T)
	if !ok {
		var zero T
		return zero, fmt.Errorf("failed to get value from context for key: %v", key)
	}
	return val, nil
}
