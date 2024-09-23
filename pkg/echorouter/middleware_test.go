package echorouter

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/zerolog"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestMiddlewareRequestID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}

	rid := MiddlewareRequestID()
	h := rid(handler)
	h(c)
	assert.Len(t, rec.Header().Get(echo.HeaderXRequestID), 36)
	_, err := uuid.Parse(rec.Header().Get(echo.HeaderXRequestID))
	assert.Nil(t, err)
}

func TestMiddlewareLogWithRequestID(t *testing.T) {
	buf := &bytes.Buffer{}
	l := zerolog.New(zerolog.ConsoleWriter{Out: buf, NoColor: true})
	log.Logger = l

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := func(c echo.Context) error {
		ctx := c.Request().Context()
		log.Ctx(ctx).Info().Msg("test")
		requestID := rec.Header().Get(echo.HeaderXRequestID)
		assert.Containsf(t, buf.String(), fmt.Sprintf(`request_id=%s`, requestID), "log message should contain request_id")
		return c.String(http.StatusOK, "test")
	}

	rid := MiddlewareRequestID()
	lrid := MiddlewareLogWithRequestID()
	h := lrid(handler)
	h = rid(h)
	h(c)
}

func TestMiddlewareRecover(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := func(c echo.Context) error {
		panic("test")
		return nil
	}

	r := MiddlewareRecover()
	h := r(handler)
	h(c)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Internal Server Error")
}
