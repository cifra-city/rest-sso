package data

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data/dbcore"
	"github.com/google/uuid"
)

type Sessions interface {
	Create(r *http.Request, userID uuid.UUID, token string, deviceName string, deviceData json.RawMessage) (dbcore.Session, error)

	GetByID(r *http.Request, id uuid.UUID) (dbcore.Session, error)
	GetSession(r *http.Request, id uuid.UUID, userID uuid.UUID) (dbcore.Session, error)
	GetSessions(r *http.Request, userID uuid.UUID) ([]dbcore.Session, error)

	GetToken(r *http.Request, id uuid.UUID, userID uuid.UUID) (string, error)
	UpdateToken(r *http.Request, id uuid.UUID, token string) error

	DeleteAll(r *http.Request, id uuid.UUID) error
	Delete(r *http.Request, id uuid.UUID, userID uuid.UUID) error
}

type sessions struct {
	queries *dbcore.Queries
}

func NewSession(queries *dbcore.Queries) Sessions {
	return &sessions{queries: queries}
}

func (s *sessions) Create(r *http.Request, userID uuid.UUID, token string, deviceName string, deviceData json.RawMessage) (dbcore.Session, error) {
	client := httpkit.GetUserAgent(r)
	IP := httpkit.GetClientIP(r)
	return s.queries.CreateSession(r.Context(), dbcore.CreateSessionParams{
		UserID:     userID,
		Token:      token,
		DeviceName: deviceName,
		Client:     client,
		Ip:         IP,
	})
}

func (s *sessions) GetByID(r *http.Request, id uuid.UUID) (dbcore.Session, error) {
	return s.queries.GetSession(r.Context(), id)
}

func (s *sessions) GetSession(r *http.Request, id uuid.UUID, userID uuid.UUID) (dbcore.Session, error) {
	return s.queries.GetUserSession(r.Context(), dbcore.GetUserSessionParams{
		ID:     id,
		UserID: userID,
	})
}

func (s *sessions) GetSessions(r *http.Request, userID uuid.UUID) ([]dbcore.Session, error) {
	return s.queries.GetSessionsByUserID(r.Context(), userID)
}

func (s *sessions) GetToken(r *http.Request, id uuid.UUID, userID uuid.UUID) (string, error) {
	return s.queries.GetSessionToken(r.Context(), dbcore.GetSessionTokenParams{
		ID:     id,
		UserID: userID,
	})
}

func (s *sessions) UpdateToken(r *http.Request, id uuid.UUID, token string) error {
	return s.queries.UpdateSessionToken(r.Context(), dbcore.UpdateSessionTokenParams{
		ID:    id,
		Token: token,
	})
}

func (s *sessions) DeleteAll(r *http.Request, id uuid.UUID) error {
	return s.queries.DeleteUserSessions(r.Context(), id)
}

func (s *sessions) Delete(r *http.Request, id uuid.UUID, userID uuid.UUID) error {
	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		return errors.New("failed to retrieve service configuration")
	}

	err = Server.TokenBin.Add(userID.String(), id.String())
	if err != nil {
		return errors.New("failed to delete device")
	}

	return s.queries.DeleteUserSession(r.Context(), dbcore.DeleteUserSessionParams{
		ID:     id,
		UserID: userID,
	})
}
