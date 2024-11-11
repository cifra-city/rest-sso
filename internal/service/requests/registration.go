package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func newDecodeError(what string, err error) error {
	return validation.Errors{
		what: fmt.Errorf("decode request %s: %w", what, err),
	}
}

func NewRegistration(r *http.Request) (req models, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	req.Data.ID = strings.ToLower(req.Data.ID)

	errs := validation.Errors{
		"data/id":                     validation.Validate(req.Data.ID, validation.Required),
		"data/type":                   validation.Validate(req.Data.Type, validation.Required, validation.In(resources.ACTIVATE_BALANCE)),
		"data/attributes/referred_by": validation.Validate(req.Data.Attributes.ReferredBy, validation.Required),
	}

	return req, errs.Filter()
}
