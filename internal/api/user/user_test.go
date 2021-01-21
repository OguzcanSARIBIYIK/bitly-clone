package user

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {

	createUser := map[string]string{
		`{"code":200,"data":"Üyelik oluşturuldu.`:        `{"username":"oguzcan17","password":"123456"}`,
		`{"code":400,"data":"Girmiş olduğunuz kullanıcı`: `{"username":"oguzcan99","password":"123456"}`,
	}

	for i, v := range createUser {

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(v))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, Register(c)) {
			fmt.Println("create : ", v)
			assert.Contains(t, rec.Body.String(), i)
		}

	}

}

func TestGetToken(t *testing.T) {

	getToken := make(map[string]string)
	getToken[`{"code":200,`] = `{"username":"oguzcan15","password":"123456"}`
	getToken[`{"code":400,"data":"Geçersiz şifre`] = `{"username":"oguzcan15","password":"1234567"}`
	getToken[`{"code":400,"data":"Kullanıcı bulunam`] = `{"username":"oguzcan188","password":"123456"}`

	for i, v := range getToken {

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(v))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, GetToken(c)) {
			fmt.Println("getToken : ", rec.Body.String())
			assert.Contains(t, rec.Body.String(), i)
		}

	}

}
