package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/pkg/jsonresp"
	"github.com/cifra-city/rest-sso/pkg/jsonresp/problems"
	"github.com/cifra-city/rest-sso/pkg/security"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	// Парсинг запроса
	req, err := requests.NewRegistration(r)
	logrus.Infof("req: %v", req)
	if err != nil {
		jsonresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// Проверка валидности пароля
	pas := req.Data.Attributes.Password
	if len(pas) < 8 || !security.HasRequiredChars(pas) {
		jsonresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// Получение данных из запроса
	em := req.Data.Attributes.Email
	username := req.Data.Id
	ctx := r.Context()

	// Проверка на существующий email
	_, err = Users(r).GetUserByEmail(ctx, em)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logrus.Errorf("error getting user by email: %v", err)
		jsonresp.RenderErr(w, problems.InternalError())
		return
	}
	if err == nil {
		jsonresp.RenderErr(w, problems.Conflict())
		return
	}

	// Проверка на существующий username
	_, err = Users(r).GetUserByUsername(ctx, username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logrus.Errorf("error getting user by username: %v", err)
		jsonresp.RenderErr(w, problems.InternalError())
		return
	}
	if err == nil {
		jsonresp.RenderErr(w, problems.Conflict())
		return
	}

	// Генерация нового UUID для пользователя
	newUuid := uuid.New()

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Data.Attributes.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("error hashing password: %v", err)
		jsonresp.RenderErr(w, problems.InternalError())
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
	user, err := Users(r).InsertUser(ctx, params)
	if err != nil {
		logrus.Errorf("error inserting user: %v", err)
		jsonresp.RenderErr(w, problems.InternalError())
		return
	}

	// Отправка успешного ответа
	logrus.Infof("user created: %v", user.Username)
	jsonresp.Render(w, map[string]string{"message": "User created successfully"})
}
