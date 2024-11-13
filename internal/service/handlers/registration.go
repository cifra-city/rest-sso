package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/pkg/cifradb"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/cifra-city/rest-sso/pkg/httpresp/problems"
	"github.com/cifra-city/rest-sso/pkg/security"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	// Парсинг запроса
	req, err := requests.NewRegistration(r)
	if err != nil {
		httpresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// Проверка валидности пароля
	pas := req.Data.Attributes.Password
	if len(pas) < 8 || !security.HasRequiredChars(pas) {
		httpresp.RenderErr(w, problems.BadRequest(errors.New("invalid password requirements"))...)
		return
	}

	// Получение данных из запроса
	em := req.Data.Attributes.Email
	username := req.Data.Id

	// Получение объекта Queries из контекста
	Q, err := cifradb.GetDBQueries(r.Context())
	if err != nil {
		logrus.Errorf("error getting db queries: %v", err)
		http.Error(w, "Database queries not found", http.StatusInternalServerError)
		return
	}

	// Проверка на существующий email
	_, err = Q.GetUserByEmail(r.Context(), em)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logrus.Errorf("error getting user by email: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}
	if err == nil {
		httpresp.RenderErr(w, problems.Conflict("this email address already exists"))
		return
	}

	// Проверка на существующий username
	_, err = Q.GetUserByUsername(r.Context(), username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logrus.Errorf("error getting user by username: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}
	if err == nil {
		httpresp.RenderErr(w, problems.Conflict("this username already exists"))
		return
	}

	// Генерация нового UUID для пользователя
	newUuid := uuid.New()

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Data.Attributes.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("error hashing password: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	// Создание параметров для вставки пользователя
	params := data.InsertUserParams{
		ID:          newUuid,
		Username:    username,
		Email:       em,
		EmailStatus: false,
		PassHash:    string(hashedPassword),
	}

	// Вставка нового пользователя в базу данных
	user, err := Q.InsertUser(r.Context(), params)
	if err != nil {
		logrus.Errorf("error inserting user: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	// Отправка успешного ответа
	logrus.Infof("user created: %v", user.Username)
	httpresp.Render(w, http.StatusCreated)
}
