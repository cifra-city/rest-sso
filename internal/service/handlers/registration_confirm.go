package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/pkg/cifractx"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/cifra-city/rest-sso/pkg/httpresp/problems"
	"github.com/cifra-city/rest-sso/pkg/sectools"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func RegistrationConfirm(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewRegistrationConfirm(r)
	if err != nil {
		httpresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	password := req.Data.Attributes.FirstPassword
	if len(password) < 8 || !sectools.HasRequiredChars(password) {
		httpresp.RenderErr(w, problems.BadRequest(errors.New("invalid password requirements"))...)
		return
	}

	if req.Data.Attributes.FirstPassword != req.Data.Attributes.SecondPassword {
		httpresp.RenderErr(w, problems.BadRequest(errors.New("passwords do not match"))...)
		return
	}

	em := req.Data.Attributes.Email
	username := req.Data.Id

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting db queries: %v", err)
		httpresp.RenderErr(w, problems.InternalError("database queries not found"))
		return
	}

	log := Server.Logger

	_, err = Server.Queries.GetUserByEmail(r.Context(), em)
	if !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("User already created: %v", err)
		httpresp.RenderErr(w, problems.Conflict("this email address already exists"))
		return
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logrus.Errorf("error getting user by email: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	_, err = Server.Queries.GetUserByUsername(r.Context(), username)
	if !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("User already created: %v", err)
		httpresp.RenderErr(w, problems.Conflict("this username already exists"))
		return
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logrus.Errorf("error getting user by username: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	if !Server.Mailman.CheckAndDeleteAccessForUser(em, "registration") {
		log.Warnf("email haven`t access: %s", em)
		httpresp.RenderErr(w, problems.Forbidden("email haven`t access"))
		return
	}

	newUuid := uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("error hashing password: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	params := data.InsertUserParams{
		ID:          newUuid,
		Username:    username,
		Email:       em,
		EmailStatus: false,
		PassHash:    string(hashedPassword),
	}

	user, err := Server.Queries.InsertUser(r.Context(), params)
	if err != nil {
		logrus.Errorf("error inserting user: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	logrus.Infof("user created: %v", user.Username)
	httpresp.Render(w, http.StatusCreated)
}
