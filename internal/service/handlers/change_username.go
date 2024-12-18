package handlers

import (
	"net/http"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/pkg/cifractx"
	"github.com/cifra-city/rest-sso/pkg/cifrajwt"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/cifra-city/rest-sso/pkg/httpresp/problems"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func ChangeUsername(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewChangeUsername(r)
	if err != nil {
		httpresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	oldPassword := req.Data.Attributes.Password
	newUsername := req.Data.Attributes.NewUsername

	service, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		httpresp.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	userID, ok := r.Context().Value(cifrajwt.UserIDKey).(uuid.UUID)
	if !ok {
		logrus.Warn("UserID not found in context")
		httpresp.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	logrus.Infof("userID: %v", userID)

	user, err := service.Queries.GetUserByID(r.Context(), userID)
	if err != nil {
		httpresp.RenderErr(w, problems.InternalError("Failed to retrieve user information"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(oldPassword))
	if err != nil {
		httpresp.RenderErr(w, problems.Unauthorized("Invalid password"))
		return
	}

	_, err = service.Queries.GetUserByUsername(r.Context(), *newUsername)
	if err == nil {
		httpresp.RenderErr(w, problems.Conflict("Username already exists"))
		return
	}

	_, err = service.Queries.UpdateUsernameByID(r.Context(), data.UpdateUsernameByIDParams{
		ID:       userID,
		Username: *newUsername,
	})
	if err != nil {
		httpresp.RenderErr(w, problems.InternalError("Failed to update username"))
		return
	}

	httpresp.Render(w, map[string]string{"message": "Username updated successfully"})
}
