package handlers

import (
	"net/http"

	"github.com/cifra-city/httpkit"
	"github.com/sirupsen/logrus"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	logrus.Debugf("user logged out")

	httpkit.Render(w, http.StatusOK)
}
