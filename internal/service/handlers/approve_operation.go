package handlers

import (
	"errors"
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/mailman"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/sirupsen/logrus"
)

type OperationType string

const (
	RESET_PASSWORD  OperationType = "reset_password"
	CHANGE_PASSWIRD OperationType = "change_password"
	CHANGE_EMAIL    OperationType = "change_email"
	REGISTRATION    OperationType = "registration"
	LOGIN           OperationType = "login"
)

func (op OperationType) IsValid() bool {
	switch op {
	case RESET_PASSWORD, REGISTRATION, LOGIN, CHANGE_PASSWIRD, CHANGE_EMAIL:
		return true
	default:
		return false
	}
}

func ApproveOperation(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewApproveOperation(r)
	if err != nil {
		logrus.Errorf("error decoding request: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	email := req.Data.Attributes.Email
	code := req.Data.Attributes.Code
	opTypeStr := req.Data.Attributes.Operation

	IP := httpkit.GetClientIP(r)
	UserAgent := httpkit.GetUserAgent(r)

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting db queries: %v", err)
		httpkit.RenderErr(w, problems.InternalError("database queries not found"))
		return
	}

	log := Server.Logger

	opType := OperationType(opTypeStr)
	if !opType.IsValid() {
		log.Errorf("invalid operation type: %s", opTypeStr)
		httpkit.RenderErr(w, problems.BadRequest(errors.New("invalid operation type"))...)
		return
	}

	err = Server.Mailman.CheckCode(email, string(opType), code, UserAgent, IP)
	if err != nil {
		if errors.Is(err, mailman.ErrNotFound) {
			log.Debugf("code not found for email: %s", email)
			httpkit.RenderErr(w, problems.NotFound("code not found"))
			return
		}
		if errors.Is(err, mailman.ErrAccessDenied) {
			log.Errorf("failed to decrypt ConfidenceCode for email: %s", email)
			httpkit.RenderErr(w, problems.InternalError("failed to decrypt ConfidenceCode"))
			return
		}
		log.Debugf("Unhadler denied err: %s", err)
		httpkit.RenderErr(w, problems.Forbidden("incorrect code"))
		return
	}

	Server.Mailman.AddAccess(email, string(opType), UserAgent, IP, 15)
	log.Debugf("code is correct add access for email: %s", email)
	httpkit.Render(w, http.StatusOK)
	return
}
