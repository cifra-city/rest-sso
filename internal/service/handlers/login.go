package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/pkg/cifractx"
	"github.com/cifra-city/rest-sso/pkg/cifrajwt"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/cifra-city/rest-sso/pkg/httpresp/problems"
	"github.com/cifra-city/rest-sso/resources"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewLogin(r)
	if err != nil {
		httpresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	email := req.Data.Attributes.Email
	pass := req.Data.Attributes.Password
	username := req.Data.Attributes.Username
	factoryId := req.Data.Attributes.FactoryId
	deviceName := req.Data.Attributes.DeviceName
	osVersion := req.Data.Attributes.OsVersion
	ipAddress := req.Data.Attributes.IpAddress

	if email == nil && username == nil {
		httpresp.RenderErr(w, problems.BadRequest(errors.New("email or username is required"))...)
		return
	}

	var user data.UsersSecret

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting server from context: %v", err)
		http.Error(w, "Service configuration not found", http.StatusInternalServerError)
		return
	}
	log := Server.Logger

	log.Debugf("email: %v, password: %s, username: %v", email, pass, username)

	if username != nil {
		user, err = Server.Queries.GetUserByUsername(r.Context(), *username)
	} else {
		user, err = Server.Queries.GetUserByEmail(r.Context(), *email)
	}
	if err != nil {
		log.Errorf("Failed to get user: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			httpresp.RenderErr(w, problems.Unauthorized())
			return
		}
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(pass))
	if err != nil {
		err = Server.Queries.InsertLoginHistory(r.Context(), data.InsertLoginHistoryParams{
			ID:        uuid.New(),
			UserID:    user.ID,
			DeviceID:  uuid.NullUUID{UUID: uuid.UUID{}, Valid: false},
			IpAddress: sql.NullString{String: ipAddress, Valid: true},
			LoginTime: time.Now().UTC(),
			Success:   false,
			FailureReason: data.NullFailureReason{
				FailureReason: data.FailureReasonInvalidPassword,
			},
		})

		log.Debugf("Incorrect password for user: %s, error: %s", user.Username, err)
		httpresp.RenderErr(w, problems.Unauthorized())
		return
	}

	tokenAccess, err := cifrajwt.GenerateJWT(user.ID, Server.Config.JWT.AccessToken.TokenLifetime, Server.Config.JWT.AccessToken.SecretKey)
	if err != nil {
		log.Errorf("error generating token access jwt: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	tokenRefresh, err := cifrajwt.GenerateJWT(user.ID, Server.Config.JWT.RefreshToken.TokenLifetime, Server.Config.JWT.RefreshToken.SecretKey)
	if err != nil {
		log.Errorf("error generating token access jwt: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	expiresAt := time.Now().UTC().Add(Server.Config.JWT.RefreshToken.TokenLifetime)

	err = Server.Queries.UpdateRefreshTokenTransaction(r.Context(), &user, factoryId, deviceName, osVersion, tokenRefresh, expiresAt, ipAddress)
	if err != nil {
		log.Errorf("error updating last used and refresh token: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}
	log.Infof("user logged in: %s", user.Username)

	httpresp.Render(w, resources.LoginResp{
		Data: resources.LoginRespData{
			Type: "login",
			Attributes: resources.LoginRespDataAttributes{
				AccessToken:  tokenAccess,
				RefreshToken: tokenRefresh,
				ExpiresIn:    int32(Server.Config.JWT.AccessToken.TokenLifetime.Seconds()),
			},
		},
	})
}
