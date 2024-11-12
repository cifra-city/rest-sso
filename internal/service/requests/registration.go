package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/cifra-city/rest-sso/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"
)

func newDecodeError(what string, err error) error {
	logrus.Info("newDecodeError")
	return validation.Errors{
		what: fmt.Errorf("decode request %s: %w", what, err),
	}
}

func NewRegistration(r *http.Request) (req resources.Registration, err error) {
	logrus.Info("NewRegistration")
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}
	logrus.Info("NewRegistration")
	req.Data.Id = strings.ToLower(req.Data.Id)

	errs := validation.Errors{
		"data/id":         validation.Validate(req.Data.Id, validation.Required),
		"data/type":       validation.Validate(req.Data.Type, validation.Required, validation.In("registrations")),
		"data/attributes": validation.Validate(req.Data.Attributes, validation.Required),
	}
	logrus.Info("NewRegistration")
	return req, errs.Filter()
}
