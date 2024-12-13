package handlers

import (
	"errors"
	"net/http"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/pkg/cifractx"
	"github.com/cifra-city/rest-sso/pkg/cifrajwt"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/cifra-city/rest-sso/pkg/httpresp/problems"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewLogin(r)
	if err != nil {
		httpresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	em := req.Data.Attributes.Email
	pas := req.Data.Attributes.Password
	usr := req.Data.Attributes.Username

	var user data.UsersSecret

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting server from context: %v", err)
		http.Error(w, "Service configuration not found", http.StatusInternalServerError)
		return
	}
	log := Server.Logger

	log.Debugf("email: %v, password: %s, username: %v", em, pas, usr)

	// Get server from context

	// work with queries from server
	if usr != nil {
		user, err = Server.Queries.GetUserByUsername(r.Context(), *usr)
		if err != nil {
			httpresp.RenderErr(w, problems.InternalError())
			return
		}
	} else if em != nil {
		user, err = Server.Queries.GetUserByEmail(r.Context(), *em)
		if err != nil {
			httpresp.RenderErr(w, problems.InternalError())
			return
		}
	} else {
		log.Infof("Bad request; email: %v, password: %s, username: %v", em, pas, usr)
		httpresp.RenderErr(w, problems.BadRequest(errors.New("email or username is required"))...)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(pas))
	if err != nil {
		log.Debugf("Incorrect password for user: %s, error: %s", user.Username, err)
		httpresp.RenderErr(w, problems.NotAllowed(errors.New("invalid password")))
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

	log.Infof("user logged in: %s", user.Username)

	httpresp.Render(w, map[string]string{"token": token})
}
