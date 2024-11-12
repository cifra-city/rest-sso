package jsonresp

import (
	"net/http"
	"strconv"

	"github.com/google/jsonapi"
)

// Render encodes a successful response in JSON API format.
func Render(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	// Serialize and write the JSON API response
	if err := jsonapi.MarshalPayload(w, data); err != nil {
		errorObject := &jsonapi.ErrorObject{
			Title:  "Internal Server Error",
			Detail: err.Error(),
			Status: strconv.Itoa(http.StatusInternalServerError),
		}
		RenderErr(w, errorObject)
	}
}
