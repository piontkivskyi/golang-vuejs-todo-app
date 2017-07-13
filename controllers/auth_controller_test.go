package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"todo-app/models/connection"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo"
)

func init() {
	connection.InitConnection()
}

func TestGetToken(t *testing.T) {
	defer connection.DB.Close()
	// Setup
	e := echo.New()
	f := make(url.Values)
	f.Set("username", "admin")
	f.Set("password", "admin")
	req := httptest.NewRequest(echo.POST, "/api/token?"+f.Encode(), nil)
	//, strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetToken(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Result())
	}
}
