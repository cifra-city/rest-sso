package problems

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
)

func Forbidden(message ...string) *jsonapi.ErrorObject {
	defaultMessage := "Forbidden"

	// Используем пользовательское сообщение, если оно передано, иначе - значение по умолчанию
	errorMessage := defaultMessage
	if len(message) > 0 && message[0] != "" {
		errorMessage = message[0]
	}

	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusForbidden),
		Status: fmt.Sprintf("%d", http.StatusForbidden),
		Detail: errorMessage,
	}
}
