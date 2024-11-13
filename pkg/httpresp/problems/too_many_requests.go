package problems

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
)

// TooManyRequests возвращает объект ошибки с HTTP статусом 429 (Too Many Requests).
func TooManyRequests(message ...string) *jsonapi.ErrorObject {
	defaultMessage := "Too many requests, please try again later"
	errorMessage := getMessageOrDefault(message, defaultMessage)

	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusTooManyRequests),
		Status: fmt.Sprintf("%d", http.StatusTooManyRequests),
		Detail: errorMessage,
	}
}
