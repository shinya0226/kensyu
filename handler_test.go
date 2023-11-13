package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string]*User{
		"shinya.yamamoto6@persol-pt.co.jp": &User{"shinya.yamamoto6@persol-pt.co.jp", "yamamo10"},
	}
	userJSON = `{"Email":"shinya.yamamoto6@persol-pt.co.jp","Password":"yamamo10"}`
)

func TestPost(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON+"\n", rec.Body.String())
	}

}
