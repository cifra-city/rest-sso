package problems

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
)

// Conflict возвращает объект ошибки с HTTP статусом 409 (Conflict).
// При отсутствии пользовательского сообщения используется сообщение по умолчанию.
func Conflict(message ...string) *jsonapi.ErrorObject {
	defaultMessage := "Resource conflict"

	// Используем пользовательское сообщение, если оно передано, иначе - значение по умолчанию
	errorMessage := defaultMessage
	if len(message) > 0 && message[0] != "" {
		errorMessage = message[0]
	}

	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusConflict),
		Status: fmt.Sprintf("%d", http.StatusConflict),
		Detail: errorMessage,
	}
}
