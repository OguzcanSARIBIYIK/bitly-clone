package user

import (
	"bitly-clone/internal/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	createUser := map[string]string{
		`{"code":201,"description":"Üyelik oluşturuldu.`: `{"username":"oguzcan17","password":"123456"}`,
		`"description":"Üyelik oluşturuldu.`:             `{"username":"oguzcan99","password":"123456"}`,
		`"description":"Girmiş olduğunuz kullanıcı`:      `{"username":"oguzcan99","password":"123456"}`,
	}
	repository.Init()
	for i, v := range createUser {

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(v))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, Register(c)) {
			assert.Contains(t, rec.Body.String(), i)
		}

	}

}

func TestGetToken(t *testing.T) {

	getToken := make(map[string]string)
	getToken[`{"code":200,`] = `{"username":"oguzcan17","password":"123456"}`
	getToken[`"data":"Geçersiz şifre`] = `{"username":"oguzcan17","password":"1234567"}`
	getToken[`"data":"Kullanıcı bulunam`] = `{"username":"oguzcan188","password":"123456"}`
	repository.Init()
	for i, v := range getToken {

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(v))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, GetToken(c)) {
			assert.Contains(t, rec.Body.String(), i)
		}

	}

}
