package handlers

import (
	"net/http"

	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/sirupsen/logrus"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	logrus.Debugf("user logged out")

	httpresp.Render(w, http.StatusOK)
}
