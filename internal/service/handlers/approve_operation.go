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
	RESET_PASSWORD OperationType = "reset_password"
	REGISTRATION   OperationType = "registration"
)

func (op OperationType) IsValid() bool {
	switch op {
	case RESET_PASSWORD, REGISTRATION:
		return true
	default:
		return false
	}
}

func ApproveOperation(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewApproveOperation(r)
	if err != nil {
		logrus.Errorf("error decoding request: %v", err)
		httpresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	email := req.Data.Attributes.Email
	code := req.Data.Attributes.Code
	opTypeStr := req.Data.Attributes.Operation

	IP := httpresp.GetClientIP(r)
	UserAgent := httpresp.GetUserAgent(r)

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting db queries: %v", err)
		httpresp.RenderErr(w, problems.InternalError("database queries not found"))
		return
	}

	log := Server.Logger

	opType := OperationType(opTypeStr)
	if !opType.IsValid() {
		log.Errorf("invalid operation type: %s", opTypeStr)
		httpresp.RenderErr(w, problems.BadRequest(errors.New("invalid operation type"))...)
		return
	}

	err = Server.Mailman.CheckCode(email, string(opType), code, UserAgent, IP)
	if err != nil {
		if errors.Is(err, errors.New("not found")) {
			log.Debugf("code not found for email: %s", email)
			httpresp.RenderErr(w, problems.NotFound("code not found"))
			return
		}
		if errors.Is(err, errors.New("access denied")) {
			log.Errorf("failed to decrypt ConfidenceCode for email: %s", email)
			httpresp.RenderErr(w, problems.InternalError("failed to decrypt ConfidenceCode"))
			return
		}
		log.Debugf("Access denied %s", email)
		httpresp.RenderErr(w, problems.Forbidden("incorrect code"))
		return
	}

	Server.Mailman.AddAccess(email, string(opType), UserAgent, IP, 15)
	log.Debugf("code is correct add access for email: %s", email)
	httpresp.Render(w, http.StatusOK)
	return
}
