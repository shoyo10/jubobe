package echorouter

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

// MiddlewareCorsConfig ...
var MiddlewareCorsConfig = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodPatch,
	},
	AllowHeaders: []string{
		"*",
	},
	ExposeHeaders: []string{
		"*",
	},
})

func requestIDGenerator() string {
	return uuid.New().String()
}

// MiddlewareRequestID middleware to set request id
func MiddlewareRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			rid := req.Header.Get(echo.HeaderXRequestID)
			if rid == "" {
				rid = requestIDGenerator()
				req.Header.Set(echo.HeaderXRequestID, rid)
			}
			res.Header().Set(echo.HeaderXRequestID, rid)
			return next(c)
		}
	}
}

// MiddlewareLogWithRequestID set request id to log field
func MiddlewareLogWithRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			rid := req.Header.Get(echo.HeaderXRequestID)
			logger := log.With().Str("request_id", rid).Logger()
			ctx := logger.WithContext(req.Context())
			c.SetRequest(req.WithContext(ctx))
			return next(c)
		}
	}
}

func MiddlewareRecover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					stack := make([]byte, 4096)
					length := runtime.Stack(stack, true)
					ctx := c.Request().Context()
					log.Ctx(ctx).Err(err).
						Str("uri", c.Request().RequestURI).
						Str("method", c.Request().Method).
						Msgf("got http runtime panic\n%+v", string(stack[:length]))

					c.Error(err)
				}
			}()
			return next(c)
		}
	}
}

var loggerConfig = middleware.LoggerConfig{
	Format: `{"level":"info","time":"${time_rfc3339_nano}","request_id":"${id}","remote_ip":"${remote_ip}",` +
		`"host":"${host}","method":"${method}","uri":"${uri}","path":"${path}","user_agent":"${user_agent}",` +
		`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
		`,"bytes_in":${bytes_in},"bytes_out":${bytes_out},"message":"http access log"}` + "\n",
}

func MiddlewareAccessLog() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(loggerConfig)
}
