package handlers

import (
	"errors"
	"net/http"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/pkg/cifractx"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/cifra-city/rest-sso/pkg/httpresp/problems"
	"github.com/sirupsen/logrus"
)

type OperationType string

const (
	FORGOT_PASSWORD OperationType = "forgot_password"
	REGISTRATION    OperationType = "registration"
)

func (op OperationType) IsValid() bool {
	switch op {
	case FORGOT_PASSWORD, REGISTRATION:
		return true
	default:
		return false
	}
}

func ApproveOperation(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewApproveOperation(r)
	if err != nil {
		httpresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	email := req.Data.Attributes.Email
	code := req.Data.Attributes.Code
	opTypeStr := req.Data.Attributes.Operation

	opType := OperationType(opTypeStr)
	if !opType.IsValid() {
		httpresp.RenderErr(w, problems.BadRequest(errors.New("invalid operation type"))...)
		return
	}

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting db queries: %v", err)
		http.Error(w, "Database queries not found", http.StatusInternalServerError)
		return
	}

	log := Server.Logger

	if Server.Mailman.CheckCode(email, string(opType), code) {
		Server.Mailman.AddAccessForUser(email, string(opType), 15)

		log.Debugf("code is correct add access for email: %s", email)
		httpresp.Render(w, http.StatusOK)
		return
	}

	log.Debugf("code is incorrect for email: %s", email)
	httpresp.Render(w, http.StatusForbidden)
}
