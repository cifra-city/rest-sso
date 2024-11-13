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
	// Получаем тело запроса
	req, err := requests.NewLogin(r)
	if err != nil {
		httpresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	em := req.Data.Attributes.Email
	pas := req.Data.Attributes.Password
	usr := req.Data.Attributes.Username

	var user data.UsersSecret

	logrus.Debugf("email: %v, password: %s, username: %v", em, pas, usr)

	// Извлечение server из контекста
	Server, err := cifractx.GetValue[*config.Server](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("error getting server from context: %v", err)
		http.Error(w, "Server configuration not found", http.StatusInternalServerError)
		return
	}

	// Работа с queries из server
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
		logrus.Infof("Bad request; email: %v, password: %s, username: %v", em, pas, usr)
		httpresp.RenderErr(w, problems.BadRequest(errors.New("email or username is required"))...)
		return
	}

	// Сравнение пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(pas))
	if err != nil {
		logrus.Debugf("Incorrect password for user: %s, error: %s", user.Username, err)
		httpresp.RenderErr(w, problems.NotAllowed(errors.New("invalid password")))
		return
	}

	// Генерация JWT токена
	token, err := cifrajwt.GenerateJWT(user.ID, Server.Config.JWT.TokenLifetime, Server.Config.JWT.SecretKey)
	if err != nil {
		logrus.Errorf("error generating jwt: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	httpresp.Render(w, map[string]string{"token": token})
}
