package echorouter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"jubobe/pkg/errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestNewEcho(t *testing.T) {
	buf := &bytes.Buffer{}
	l := zerolog.New(zerolog.ConsoleWriter{Out: buf, NoColor: true})
	log.Logger = l

	e := NewEcho(&Config{})
	e.GET("/testlog", func(c echo.Context) error {
		log.Ctx(c.Request().Context()).Info().Msg("test")
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/panic", func(c echo.Context) error {
		panic("test panic")
		return nil
	})

	rec := request(http.MethodGet, "/ping", e)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "pong", rec.Body.String())
	assert.Len(t, rec.Header().Get(echo.HeaderXRequestID), 36)

	rec = request(http.MethodGet, "/testlog", e)
	requestID := rec.Header().Get(echo.HeaderXRequestID)
	assert.Containsf(t, buf.String(), fmt.Sprintf(`request_id=%s`, requestID), "log message should contain request_id")

	rec = request(http.MethodGet, "/panic", e)
	requestID = rec.Header().Get(echo.HeaderXRequestID)
	assert.Containsf(t, buf.String(), fmt.Sprintf(`request_id=%s`, requestID), "log message should contain request_id")
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	var result errors.HTTPError
	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.Equal(t, errors.ErrInternalServerError.Code, result.Code)
	assert.Equal(t, errors.ErrInternalServerError.Message, result.Message)

	rec = request(http.MethodGet, "/notfoundpage", e)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t, "page not found", rec.Body.String())

	rec = request(http.MethodGet, "/debug/pprof", e)
	assert.Contains(t, rec.Body.String(), "/debug/pprof")
}

func request(method, path string, e *echo.Echo) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	return rec
}

func TestFxNewEcho(t *testing.T) {
	var e *echo.Echo
	app := fx.New(
		fx.Supply(&Config{}),
		fx.Provide(FxNewEcho),
		fx.Populate(&e),
	)
	err := app.Start(context.Background())
	assert.Nil(t, err)

	defer app.Stop(context.Background())

	rec := request(http.MethodGet, "/ping", e)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "pong", rec.Body.String())
}
