package jsonresp

import (
	"log"
	"net/http"
	"strconv"

	"github.com/google/jsonapi"
)

func RenderErr(w http.ResponseWriter, errs ...*jsonapi.ErrorObject) {
	if len(errs) == 0 {
		log.Println("No errors provided to RenderErr")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	status, err := strconv.Atoi(errs[0].Status)
	if err != nil {
		log.Printf("Invalid status in error object: %v", errs[0].Status)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(status)

	if err := jsonapi.MarshalErrors(w, errs); err != nil {
		log.Printf("Failed to marshal errors: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
