package httpresp

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/jsonapi"
	"github.com/pkg/errors"
)

// Render encodes a successful response in JSON API format.
func Render(w http.ResponseWriter, res interface{}) {
	w.Header().Set("content-type", jsonapi.MediaType)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(errors.Wrap(err, "failed to render response"))
	}
}

func RenderErr(w http.ResponseWriter, errs ...*jsonapi.ErrorObject) {
	if len(errs) == 0 {
		log.Println("No errors provided to RenderErr")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Парсинг статуса из первой ошибки
	status, err := strconv.Atoi(errs[0].Status)
	if err != nil {
		log.Printf("Invalid status in error object: %v", errs[0].Status)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Установка заголовков и статуса ответа
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(status)

	// Обработка маршализации и вывод ошибок JSON API
	if err := jsonapi.MarshalErrors(w, errs); err != nil {
		log.Printf("Failed to marshal errors: %v", err)
		// Если произошла ошибка маршализации, ничего не делаем, чтобы избежать повторного WriteHeader
	}
}
