package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/rest-sso/internal/sectools"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func RegistrationComplete(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewRegistrationComplete(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	password := req.Data.Attributes.FirstPassword
	em := req.Data.Attributes.Email
	username := req.Data.Attributes.Username

	IP := httpkit.GetClientIP(r)
	UserAgent := httpkit.GetUserAgent(r)

	if len(password) < 8 || !sectools.HasRequiredChars(password) {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("invalid password requirements"))...)
		return
	}

	if req.Data.Attributes.FirstPassword != req.Data.Attributes.SecondPassword {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("passwords do not match"))...)
		return
	}

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting db queries: %v", err)
		httpkit.RenderErr(w, problems.InternalError("database queries not found"))
		return
	}

	log := Server.Logger

	_, err = Server.Queries.GetUserByEmail(r.Context(), em)
	if !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("User already created: %v", err)
		httpkit.RenderErr(w, problems.Conflict("this email address already exists"))
		return
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logrus.Errorf("error getting user by email: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	_, err = Server.Queries.GetUserByUsername(r.Context(), *username)
	if !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("User already created: %v", err)
		httpkit.RenderErr(w, problems.Conflict("this username already exists"))
		return
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logrus.Errorf("error getting user by username: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	err = Server.Mailman.CheckAccess(em, "registration", UserAgent, IP)
	if err != nil {
		if errors.Is(err, errors.New("not found")) {
			log.Warnf("email haven`t access: %s", em)
			httpkit.RenderErr(w, problems.NotFound("email haven`t access"))
			return
		}
		if errors.Is(err, errors.New("access denied")) {
			log.Warnf("failed to decrypt ConfidenceCode for email: %s", em)
			httpkit.RenderErr(w, problems.Forbidden("failed to decrypt ConfidenceCode"))
			return
		}
		log.Warnf("Access denied %s, %s %s", err, IP, UserAgent)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	newUuid := uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("error hashing password: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	params := data.InsertUserParams{
		ID:       newUuid,
		Username: *username,
		Email:    em,
		PassHash: string(hashedPassword),
	}

	user, err := Server.Queries.InsertUser(r.Context(), params)
	if err != nil {
		logrus.Errorf("error inserting user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	logrus.Infof("user created: %v", user.Username)
	httpkit.Render(w, http.StatusCreated)
}
