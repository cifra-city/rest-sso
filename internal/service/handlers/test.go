package handlers

import (
	"net/http"

	"github.com/cifra-city/rest-sso/pkg/cifrajwt"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/sirupsen/logrus"
)

func Testw(w http.ResponseWriter, r *http.Request) {
	logrus.Debugf("user test")

	token, err := cifrajwt.ExtractToken(r.Context())
	if err != nil {
		http.Error(w, "Failed to extract token: "+err.Error(), http.StatusBadRequest)
		return
	}

	logrus.Infof("Extracted Token: %s", token)

	httpresp.Render(w, http.StatusOK)
}
